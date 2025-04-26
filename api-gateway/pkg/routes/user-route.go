package routes

import (
	"api-gateway/pkg/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userService *services.UserService
}

func NewUserRoutes(userService *services.UserService)*UserRoutes{
	return &UserRoutes{
		userService: userService,
	}
}

func(ur *UserRoutes)RegisterRoute(c *gin.Context){
	var req struct {
        Email     string `json:"email" binding:"required,email"`
        Password  string `json:"password" binding:"required"`
        FirstName string `json:"first_name" binding:"required"`
        LastName  string `json:"last_name" binding:"required"`
        Role      string `json:"role" binding:"required,oneof=ADMIN CUSTOMER SELLER"`
    }
		if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Bad request",
			"error":err.Error(),
		})
		return
	}
	res,err:=ur.userService.RegisterService(req.Email,req.Password,req.FirstName,req.LastName,req.Role)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message":"user created successfully",
		"data":res,
	})
}
func(ur *UserRoutes)LoginRoute(c *gin.Context){
	var req struct{
		Email string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err:=c.ShouldBindJSON(&req);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Bad request",
			"error":err.Error(),
		})
		return
	}
	res,err:=ur.userService.LoginService(req.Email,req.Password)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"Internal Server Error",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"data":res,
	})
}
func(ur *UserRoutes)GetUserByID(c *gin.Context){
	idPraram:=c.Param("id")
	res,err:=ur.userService.GetUserByIdService(idPraram)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"user retreived successfully",
		"data":res,
	})
}