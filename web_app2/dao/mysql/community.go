package mysql

import "web_app2/models"

func GetCommunityList() (communityList []*models.Community, err error) {
	result := db.Table("community").Select("community_id", "community_name", "introduction").Find(&communityList)
	if result.Error != nil {
		err = result.Error
	}
	return communityList, err
}

func GetCommunityByID(id int64) (community *models.CommunityDetail, err error) {
	result := db.Table("community").Where("community_id=?", id).Find(&community)
	if result.Error != nil {
		err = result.Error
	}
	return community, err
}
