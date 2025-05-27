package logic

import (
	"errors"
	"fmt"
	"time"
	"web_app2/dao/mysql"
	"web_app2/models"
	"web_app2/pkg/jwt"
	"web_app2/pkg/snowflake"
)

//存放业务逻辑代码

func SignUp(p *models.ParamSighUp) (err error) {
	// 判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		fmt.Println("mysql.CheckUserExist err:", err)
		return err
	}
	if err == mysql.ErrorUserExist {
		return errors.New("用户已存在")
	}

	// 生成UID
	userID := snowflake.GenID()
	//构造user实例
	u := &models.User{
		UserId:     userID,
		UserName:   p.Username,
		Password:   p.Password,
		CreateTime: time.Now(),
	}
	// 保存进数据库

	return mysql.InsertUser(u)

}

func SignIn(p *models.ParamSignIn) (user *models.User, err error) {
	user = &models.User{
		UserName: p.Username,
		Password: p.Password,
	}

	//判断用户是否存在,  传递的是指针，因此user变量改变了
	if err := mysql.FindUser(user); err != nil {
		return nil, err
	}
	// 生成jwt
	atoken, rtoken, err := jwt.GenerateToken(user.UserId)
	if err != nil {
		return nil, err
	}
	user.Atoken = atoken
	user.Rtoken = rtoken
	return user, nil
}
