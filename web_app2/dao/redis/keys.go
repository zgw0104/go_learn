package redis

// redis key

const (
	KeyPrefix              = "catcat"
	KeyPostTimeZset        = "post:time"   //zset：帖子以发帖时间为分数
	KeyPostScoreZset       = "post:score"  //zset：帖子以投票为分数
	KeyPostVotedZsetPrefix = "post:voted:" //记录用户及投票类型;参数是 postid
	KeyPostCommunitySet    = "community:"  //保存每个分区下帖子的id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
