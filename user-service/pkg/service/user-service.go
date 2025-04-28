package service

import (
	"log"
	"time"
	pb "user-service/gen/user"
	"user-service/pkg/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string)(string,error){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return string(bytes),err
}
func checkPasswordHash(password,hash string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err==nil
}
func genertaeJWT(userID string,role string)(string,error){
	jwtSecret:=[]byte("jwt-secret")
	claims:=jwt.MapClaims{
		"user_id":userID,
		"user_role":role,
		"exp":time.Now().Add(time.Hour*24).Unix(),
	}
	token:=jwt.NewWithClaims(jwt.SigningMethodHS512,claims)
	return token.SignedString(jwtSecret) 
}
type UserService struct {
	userRepo *repository.UserRepository
}
func NewUserService(userRepo *repository.UserRepository)*UserService{
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService)Register(email,password,first_name,last_name,role string)(*pb.CreateUserResponse,error){
	hash,err:=hashPassword(password)
	if err!=nil{
		return nil,err
	}
	return u.userRepo.CreateUser(email,hash,first_name,last_name,role)
}
func (u *UserService)Login(email,password string)(*pb.AuthenticateUserResponse,error){ 
	user,err:=u.userRepo.GetUserByEmail(email)
	if err!=nil{
		log.Println("error in get user by email : ",err)
	}
	if !checkPasswordHash(password,user.Password){
		return nil,err
	}
	token,err:=genertaeJWT(user.Id,user.Role.String())
	if err!=nil{
		return nil,err
	}
	return &pb.AuthenticateUserResponse{
		Token: token,
		User: user,
		Message: "Success",
	},nil
}
func (u *UserService)GetUserByID(id int)(*pb.GetUserResponse,error){
	return u.userRepo.GetUserByID(id)
}