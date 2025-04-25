package handlers

import (
	"context"
	"fmt"
	"strconv"
	pb "user-service/gen/user"
	"user-service/pkg/service"
)
type UserHandlers struct{
	userService *service.UserService
	pb.UnimplementedUserServiceServer
}
func NewUserHandlers(userSvc *service.UserService)*UserHandlers{
	return &UserHandlers{
		userService: userSvc,
	}
}
func (h *UserHandlers)CreateUser(ctx context.Context,req *pb.CreateUserRequest)(*pb.CreateUserResponse,error){
	res,err:=h.userService.Register(req.Email,req.Password,req.FirstName,req.LastName,req.Role.String())
	if err!=nil{
		return nil,fmt.Errorf("error in CreateUser Handler: %v",err)
	}
	 return res,nil
}
func (h *UserHandlers)GetUser(ctx context.Context,req *pb.GetUserRequest)(*pb.GetUserResponse,error){
	id,err:=strconv.Atoi(req.UserId)
	if err!=nil{
		return nil,fmt.Errorf("error in converting stirng in GetUser Handler: %v",err)
	}
	res,err:= h.userService.GetUserByID(id)
	if err!=nil{
		return nil,fmt.Errorf("error in GetUser Handler: %v",err)
	}
	return res,nil
}
func (h *UserHandlers) Login(ctx context.Context,req *pb.AuthenticateUserRequest)(*pb.AuthenticateUserResponse,error){
	res,err:=h.userService.Login(req.Email,req.Password)
	if err!=nil{
		return nil,fmt.Errorf("error in Login Handler: %v",err)
	}
	return res,nil
}