package routes

import (
	"api-gateway/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProductsRoute (c *gin.Context){
	pageParam:=c.DefaultQuery("page","1")
	limitParam:=c.DefaultQuery("limit","5")

	page, err1 := strconv.Atoi(pageParam)
	limit, err2 := strconv.Atoi(limitParam)

	if err1 != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters. 'page' and 'limit' must be integers.",
		})
		return
	}
	res,err:=services.GetAllProductService(int32(page),int32(limit))
	if res ==nil ||len(res.Products)==0{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"No product found",
		})
		return
	}
	fmt.Println("Response from Get All products",res)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Products Retrieved successfully",
		"data":res,
	})
}
func GetProductByIdRoute(c *gin.Context){
	IdParam:=c.Param("id")
	Id,err:=strconv.Atoi(IdParam)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Bad Request",
		})
		return
	}
	res,err:=services.GetProductByIdService(Id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}
	if res.Product ==nil{
		c.JSON(http.StatusNotFound,gin.H{
			"message":"No product found",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Product Retrieved successfully",
		"data":res,
	})
}