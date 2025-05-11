package logic

import (
	"errors"
	"fmt"
	"time"
	"web_app2/dao/mysql"
	"web_app2/models"
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

func SignIn(p *models.ParamSignIn) (err error) {
	user := &models.User{
		UserName: p.Username,
		Password: p.Password,
	}

	fmt.Println("2) user:= &U{}", *user)
	//判断用户是否存在
	if err := mysql.FindUser(user); err != nil {
		return err
	}

	return
}
