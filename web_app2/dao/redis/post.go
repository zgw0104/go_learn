package redis

import (
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"strconv"
	"time"
	"web_app2/models"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	// 3. ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return Rdb.ZRevRange(key, start, end).Result()
}

func CreatePost(postID, communityId int64) error {
	// 帖子时间
	pipeline := Rdb.TxPipeline()
	zap.L().Debug("time_now", zap.Float64("time_now", float64(time.Now().Unix())))
	pipeline.ZAdd(getRedisKey(KeyPostTimeZset), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(getRedisKey(KeyPostScoreZset), redis.Z{
		Score:  0,
		Member: postID,
	})

	// 把帖子id加到社区set
	pipeline.SAdd(getRedisKey(KeyPostCommunitySet+strconv.Itoa(int(communityId))), postID)

	_, err := pipeline.Exec()
	return err
}

// 从redis 获取id
func GetPostsIdByOrder(p *models.ParamPostList) ([]string, error) {

	key := getRedisKey(KeyPostTimeZset)
	if p.Order == "score" {
		key = getRedisKey(KeyPostScoreZset)
	}

	return getIDsFormKey(key, int64(p.Page), int64(p.Pagesize))
}

func GetPostVoteData(ids []string) (data []int64, err error) {
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZsetPrefix + id)
	//
	//	v1 := Rdb.ZCount(key, "1", "1").Val() //统计当前key下的值区间的数量
	//	data = append(data, v1)
	//}
	pipeline := Rdb.TxPipeline()

	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZsetPrefix + id)
		pipeline.ZCount(key, "1", "1")
	}
	cmds, err := pipeline.Exec()
	if err != nil {
		return nil, err

	}
	data = make([]int64, 0, len(cmds))
	for _, cmd := range cmds {
		v := cmd.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return data, nil
}

//按社区 根据ids查找数据
func GetCommunityPostsIdByOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZset)
	if p.Order == "score" {
		orderKey = getRedisKey(KeyPostScoreZset)
	}
	//使用interstore把分区帖子set 和按分数排序的zset 生成新的zset
	//针对新的zset 按之前的逻辑取数据
	//社区key
	ckey := getRedisKey(KeyPostCommunitySet + strconv.Itoa(int(p.CommunityID)))
	//利用缓存key减少zinterstore执行的次数
	key := orderKey + strconv.Itoa(int(p.CommunityID))
	if Rdb.Exists(key).Val() < 1 {
		//不存在 需要计算
		pipeline := Rdb.TxPipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, ckey, orderKey)
		pipeline.Expire(key, time.Second*60)
		pipeline.Exec()
	}
	return getIDsFormKey(key, p.Page, p.Pagesize)
}
