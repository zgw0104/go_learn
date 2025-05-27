package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app2/logic"
	"web_app2/models"
)

// 创建帖子
func CreatePostHandler(c *gin.Context) {
	//1  获取参数及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Debug("c.ShouldBindJSON(p)", zap.Any("err", err))
		zap.L().Error("Post with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	// 从当前 c 中取到发请求用户的id
	userId, err := GetCurrentUserId(c)
	if err != nil {
		Response(c, CodeInvalidToken)
		return
	}
	//2  创建帖子
	p.AuthorId = userId
	err = logic.CreatePost(p)
	if err != nil {
		zap.L().Error("Post with invalid param", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}

	//3  返回响应
	ResponseSuccess(c, nil)
}

// 获取帖子详情
func GetPostDetailHandler(c *gin.Context) {
	// 获取帖子id
	postId := c.Param("id")
	id, err := strconv.ParseInt(postId, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		Response(c, CodeInvalidParam)
		return
	}

	//根据id 查询帖子内容
	post, err := logic.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail failed", zap.Error(err))
		ResponseErrorWithMsg(c, CodeInvalidParam, err.Error())
		return
	}
	ResponseSuccess(c, post)

}

func GetPostListHandler(c *gin.Context) {

	ipage, ipagesize := GetPageInfo(c)
	postList, err := logic.GetPostList(ipage, ipagesize)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		Response(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postList)
}

// 根据前端传来的参数
// 按分数  or  创建时间 获取帖子列表

func GetPostListHandler2(c *gin.Context) {

	// 1、 获取参数
	// 2、 去redis取id列表
	// 3、 根据id去db查帖子详细信息

	// Get 请求参数： /api/v1/posts2?page=1&pagesize=5&order=time(or score)
	p := &models.ParamPostList{
		Page:     1,
		Pagesize: 10,
		Order:    "time",
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("ShouldBindQuery failed", zap.Error(err))
		Response(c, CodeInvalidParam)
		return
	}

	//ipage, ipagesize, order := GetPageInfo2(c)
	postList, err := logic.GetPostList2(p)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		Response(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postList)
}

// 根据社区去查询帖子列表

// GetPostListHandler2 升级版帖子列表接口
// @Summary 升级版帖子列表接口
// @Description 可按社区按时间或分数排序查询帖子列表接口
// @Tags 帖子相关接口
// @Accept application/json
// @Produce application/json
// @Param Authorization header string false "Bearer 用户令牌"
// @Param object query models.ParamPostList false "查询参数"
// @Security ApiKeyAuth
// @Success 200 {object} _ResponsePostList
// @Router /posts2 [get]
func GetCommunityPostListHandler(c *gin.Context) {

	p := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:     1,
			Pagesize: 10,
			Order:    "time",
		},
	}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("ShouldBindQuery failed", zap.Error(err))
		Response(c, CodeInvalidParam)
		return
	}

	postList, err := logic.GetCommunityPostList(p)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		Response(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postList)
}
