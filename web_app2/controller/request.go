package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const ContextUserIdKey = "user_id"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUserId(c *gin.Context) (userId int64, err error) {
	uid, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

func GetPageInfo(c *gin.Context) (int, int) {
	//获取分页参数
	pageStr := c.Query("page")
	pagesizeStr := c.Query("pagesize")

	var (
		page     int64
		pagesize int64
		err      error
	)

	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 0
	}
	pagesize, err = strconv.ParseInt(pagesizeStr, 10, 64)
	if err != nil {
		pagesize = 10
	}
	ipage := int(page)
	ipagesize := int(pagesize)
	return ipage, ipagesize
}
