package logic

import (
	"go.uber.org/zap"
	"strconv"
	"web_app2/dao/redis"
	"web_app2/models"
)

func PostVote(userID int64, p *models.VoteData) error {
	zap.L().Debug("VoteForPost", zap.Int64("user_id", userID), zap.String("post_id", p.PostID), zap.Int("direction", int(p.Direction)))
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
