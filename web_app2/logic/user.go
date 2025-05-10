package logic

import (
	"web_app2/dao/mysql"
	"web_app2/models"
	"web_app2/pkg/snowflake"
)

//存放业务逻辑代码

func SignUp(p *models.ParamSighUp) {
	// 判断用户存不存在
	mysql.QueryUserByUserName()

	// 生成UID
	snowflake.GenID()

	// 保存进数据库
	mysql.InsertUser()
}
