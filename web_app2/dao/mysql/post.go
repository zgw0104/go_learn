package mysql

import (
	"strings"
	"web_app2/dao/redis"
	"web_app2/models"
)

func CreatePost(p *models.Post) error {

	result := db.Table("post").Create(&p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetPostDetail(id int64) (postDetail *models.ApiPostDetail, err error) {
	post := &models.Post{}
	//result := db.Table("post").Where("post_id=?", id).First(&post)
	result := db.Raw("select * from post where post_id = (?)", id).Scan(&post)
	if result.Error != nil {
		return nil, result.Error
	}

	author := &models.User{}
	//authorResult := db.Table("user").Where("user_id=?", post.AuthorId).First(&author)
	authorResult := db.Raw("select * from user where user_id = (?)", post.AuthorId).Scan(&author)
	if authorResult.Error != nil {
		return nil, result.Error
	}

	community := &models.Community{}
	//communityResult := db.Table("community").Where("community_id=?", post.Community_id).First(&community)
	communityResult := db.Raw("select * from community where community_id = (?)", post.Community_id).Scan(&community)
	if communityResult.Error != nil {
		return nil, communityResult.Error
	}

	postDetail = &models.ApiPostDetail{
		AuthorName: author.UserName,
		Post:       post,
		Community:  community,
	}

	return postDetail, nil
}

func GetPostList(page, pagesize int) (postList []*models.ApiPostDetail, err error) {
	posts := make([]*models.Post, 0, 5)

	result := db.Table("post").Limit(pagesize).Offset((page - 1) * pagesize).Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	postList = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		author := &models.User{}
		authorResult := db.Table("user").Where("user_id=?", post.AuthorId).First(&author)
		if authorResult.Error != nil {
			return nil, result.Error
		}

		community := &models.Community{}
		communityResult := db.Table("community").Where("community_id=?", post.Community_id).First(&community)
		if communityResult.Error != nil {
			return nil, communityResult.Error
		}

		postDetail := &models.ApiPostDetail{
			AuthorName: author.UserName,
			Post:       post,
			Community:  community,
		}
		postList = append(postList, postDetail)

	}

	return postList, nil
}

// 根据给定id列表查询数据
func GetPostList2(postIds []string) (postList []*models.ApiPostDetail, err error) {
	posts := make([]*models.Post, 0, len(postIds))
	// gorm.Expr: 将sql语句作为参数执行，  Order(FIND_IN_SET(XX,?)) 根据给定的字符串序列中的元素(以逗号分割)顺序输出
	//result := db.Table("post").Where("post_id in ?", postIds).Order(gorm.Expr("FIND_IN_SET(post_id, ?) DESC", strings.Join(postIds, ","))).Find(&posts)
	result := db.Raw("select * from post where post_id in (?) ORDER BY find_in_set(post_id, ?)", postIds, strings.Join(postIds, ",")).Scan(&posts)
	if result.Error != nil {
		return nil, result.Error
	}

	data, err := redis.GetPostVoteData(postIds)

	postList = make([]*models.ApiPostDetail, 0, len(posts))
	for id, post := range posts {
		author := &models.User{}
		authorResult := db.Table("user").Where("user_id=?", post.AuthorId).First(&author)
		if authorResult.Error != nil {
			return nil, authorResult.Error
		}
		community := &models.Community{}
		communityResult := db.Table("community").Where("community_id=?", post.Community_id).First(&community)
		if communityResult.Error != nil {
			return nil, communityResult.Error
		}

		postDetail := &models.ApiPostDetail{
			AuthorName: author.UserName,
			Score:      data[id],
			Post:       post,
			Community:  community,
		}
		postList = append(postList, postDetail)
	}

	return postList, nil

}
