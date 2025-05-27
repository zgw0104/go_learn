package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web_app2/logic"
)

func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区 (id, name, )以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		Response(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}

// 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	//1、获取社区id
	communityId := c.Param("id")
	id, err := strconv.ParseInt(communityId, 10, 64)
	if err != nil {
		zap.L().Error("strconv.ParseInt failed", zap.Error(err))
		Response(c, CodeInvalidParam)
	}
	// 查询到所有的社区 (id, name, )以列表的形式返回
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		Response(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
