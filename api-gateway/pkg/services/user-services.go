package services

import (
	pb "api-gateway/gen/user"
	"context"
	"fmt"
	"time"
)
type UserService struct {
	client pb.UserServiceClient
}
func NewUserService(client pb.UserServiceClient)*UserService{
	return &UserService{
		client: client,
	}
}

func (s *UserService)RegisterService(email,password,first_name,last_name,role string)(*pb.CreateUserResponse,error){
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var pbRole pb.Role
	switch role{
	case "ADMIN":
		pbRole=pb.Role_ADMIN
	case "CUSTOMER":
		pbRole=pb.Role_CUSTOMER
	case "SELLER":
		pbRole=pb.Role_SELLER
	default:
        return nil, fmt.Errorf("invalid role: %s", role)
	}
	res,err:=s.client.CreateUser(ctx,&pb.CreateUserRequest{
		Email: email,
		Password: password,
		FirstName: first_name,
		LastName: last_name,
		Role: pbRole,
	})
	if err!=nil{
		return nil,err
	}
	return res,nil
}
func (s *UserService)LoginService(email,pass string)(*pb.AuthenticateUserResponse,error){
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	res,err:=s.client.Login(ctx,&pb.AuthenticateUserRequest{
		Email: email,
		Password: pass,
	})
	if err!=nil{
		return nil,err
	}
	return res,err
}
func (s *UserService)GetUserByIdService(id string)(*pb.GetUserResponse,error){

	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	res,err:=s.client.GetUser(ctx,&pb.GetUserRequest{
		UserId: id,
	})
	if err!=nil{
		return nil,err
	}
	return res,err
}