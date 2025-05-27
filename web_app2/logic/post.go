package logic

import (
	"go.uber.org/zap"
	"time"
	"web_app2/dao/mysql"
	"web_app2/dao/redis"
	"web_app2/models"
	"web_app2/pkg/snowflake"
)

func CreatePost(p *models.Post) error {
	//1.生成postID
	p.PostId = snowflake.GenID()
	p.CreateTime = time.Now()
	p.UpdateTime = time.Now()

	//2.入库
	err := mysql.CreatePost(p)
	if err != nil {
		return err
	}
	err = redis.CreatePost(p.PostId, p.Community_id)
	return err

}

func GetPostDetailByID(id int64) (*models.ApiPostDetail, error) {
	post, err := mysql.GetPostDetail(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetail", zap.Error(err))
		return nil, err
	}

	return post, nil
}

func GetPostList(page, pagesize int) ([]*models.ApiPostDetail, error) {

	postList, err := mysql.GetPostList(page, pagesize)
	if err != nil {
		zap.L().Error("mysql.GetPostList", zap.Error(err))
		return nil, err
	}

	return postList, nil
}

func GetPostList2(p *models.ParamPostList) ([]*models.ApiPostDetail, error) {
	// 去redis 查询id列表
	postsId, err := redis.GetPostsIdByOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostsIdByOrder failed", zap.Error(err))
		return nil, err
	}
	if len(postsId) == 0 {
		zap.L().Warn("redis.GetPostsIdByOrder return 0")
		return nil, err
	}
	// 提前查询每篇帖子的投票数
	// 返回的数据需要根据我给定的数据返回
	postList, err := mysql.GetPostList2(postsId)
	if err != nil {
		zap.L().Error("mysql.GetPostList", zap.Error(err))
		return nil, err
	}

	return postList, nil
}

func GetCommunityPostList(p *models.ParamCommunityPostList) ([]*models.ApiPostDetail, error) {
	// 去redis 查询id列表
	postsId, err := redis.GetCommunityPostsIdByOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostsIdByOrder failed", zap.Error(err))
		return nil, err
	}
	if len(postsId) == 0 {
		zap.L().Warn("redis.GetPostsIdByOrder return 0")
		return nil, err
	}
	// 提前查询每篇帖子的投票数
	// 返回的数据需要根据我给定的数据返回
	postList, err := mysql.GetPostList2(postsId)
	if err != nil {
		zap.L().Error("mysql.GetPostList", zap.Error(err))
		return nil, err
	}

	return postList, nil
}
