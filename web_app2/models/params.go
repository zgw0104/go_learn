package models

//定义请求参数的结构体

// 用户注册请求参数
type ParamSighUp struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RePasswd string `json:"re_passwd" binding:"required,eqfield=Password"`
}

// 用户登录请求参数
type ParamSignIn struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// 用户投票参数
type VoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction"` //赞成(1)or反对(-1) 取消(0)

}

//获取帖子列表参数
type ParamPostList struct {
	Page        int64  `json:"page" form:"page"`
	Pagesize    int64  `json:"pagesize" form:"pagesize"`
	Order       string `json:"order" form:"order"`
	CommunityID int64  `json:"community_id" form:"community_id"` //form  将query表单绑定(shouldBind)
}

//根据社区获取帖子列表
type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"` //form  将query表单绑定(shouldBind)
}
