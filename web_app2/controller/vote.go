package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web_app2/logic"
	"web_app2/models"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验:
	p := new(models.VoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors) //类型断言
		if !ok {
			Response(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
	}

	userID, err := GetCurrentUserId(c)
	if err != nil {
		Response(c, CodeNeedLogin)
		return
	}
	if err := logic.PostVote(userID, p); err != nil {
		zap.L().Error("logic.PostVote failed", zap.Error(err))
		Response(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}
