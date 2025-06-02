package service

import (
	"context"
	"kratos_git/models"
	"kratos_git/pkg/jwt"

	pb "kratos_git/api/git"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	ub := new(models.UserBasic)
	err := models.DB.Table("user_basic").Where("username=? and password=?", req.Username, req.Password).First(ub).Error
	if err != nil {
		return nil, err
	}
	atoken, _, err := jwt.GenerateToken(ub.Identity)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Token: atoken,
	}, nil
}
