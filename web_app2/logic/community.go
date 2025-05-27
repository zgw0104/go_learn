package logic

import (
	"go.uber.org/zap"
	"web_app2/dao/mysql"
	"web_app2/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查找所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	community, err := mysql.GetCommunityByID(id)
	if err != nil {
		zap.L().Error("mysql.GetComm unityByID failed", zap.Error(err))
		return nil, err
	}

	return community, nil
}
