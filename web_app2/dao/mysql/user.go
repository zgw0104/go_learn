package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"go.uber.org/zap"
	"time"
	"web_app2/models"
)

const secret = "zgw0104"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
)

// 把每一步数据库操作封装成函数等待logic层根据业务需求调用

func CheckUserExist(username string) (err error) {
	var users []models.User
	var count int64
	result := db.Table("user").Where("username = ?", username).Find(&users).Count(&count)
	if result.Error != nil {
		// 数据库查询错误
		return result.Error
	}

	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func InsertUser(user *models.User) (err error) {

	user.Password = encryptPwd(user.Password)
	user.UpdateTime = time.Now()
	// 执行sql语句入库
	result := db.Table("user").Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUser(user *models.User) (err error) {
	tmp := encryptPwd(user.Password)
	result := db.Table("user").Where("username = ? ", user.UserName).Find(&user)
	if result.Error != nil {
		return ErrorUserNotExist
	}
	if tmp != user.Password {
		return ErrorInvalidPassword
	}
	return nil
}

func encryptPwd(pwd string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(pwd)))
}

func FindUserByID(id int64) (user *models.User, err error) {
	result := db.Table("user").Where("user_id = ?", id).First(&user)
	if result.Error != nil {
		zap.L().Error("User not exist", zap.Error(result.Error))
		return nil, result.Error
	}
	return user, nil
}
