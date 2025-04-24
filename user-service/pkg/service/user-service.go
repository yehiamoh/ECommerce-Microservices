package service

import "user-service/pkg/repository"

type UserService struct {
	userRepo *repository.UserRepository
}
func NewUserService(userRepo *repository.UserRepository)*UserService{
	return &UserService{
		userRepo: userRepo,
	}
}
