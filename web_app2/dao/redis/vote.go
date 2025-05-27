package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"math"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432
)

var (
	ErrorTimeExpire = errors.New("超出投票时间")
	ErrVoteRepeated = errors.New("重复投票")
)

// 简化版投票分数
// 投一票+432分   86400/200 -> 需要200个赞成 才能让帖子续一天在首页

/* 投票的几种情况
 direction=1:
	1. 之前没投过票，现在投赞成	+432	ov=0	dir=1
	2. 之前投反对，现在投赞成	+864	ov=-1	dir=1
 direction=0:
	1.之前投反对，现在投取消	+432	ov=-1	dir=0
	2.之前投赞成，现在投取消	-432	ov=1	dir=0
 direction=-1:
	1.之前没投票，现在投反对	-432	ov=0	dir=-1
	2.之前投赞成，现在投反对	-864	ov=1	dir=-1

投票的限制：
每个帖子发表之日起一个星期之内允许用户投票，超过一个星期就不准投票了
	1.到期之后将redis中保存的赞成票及反对票存储到mysql
	2.到期后删除 key

*/

func VoteForPost(userID, postID string, value float64) error {
	// 判断投票限制
	// 去redis取帖子发布时间
	postTime := Rdb.ZScore(getRedisKey(KeyPostTimeZset), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		zap.L().Debug("time", zap.Float64("time", float64(time.Now().Unix())-postTime))
		return ErrorTimeExpire
	}
	// 更新分数
	// 查当前用户给当前帖子的投票纪录
	ov := Rdb.ZScore(getRedisKey(KeyPostVotedZsetPrefix+postID), userID).Val() //post:vote:3132131231245.用户a = ov
	//zap.L().Debug("vote", zap.Float64("vote", float64(ov)))
	//zap.L().Debug("value", zap.Float64("value", float64(value)))

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := Rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZset), op*diff*scorePerVote, postID)

	// 3. 记录用户为该贴子投票的数据
	if value == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZsetPrefix+postID), userID)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZsetPrefix+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
