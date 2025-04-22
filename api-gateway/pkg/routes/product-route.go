package routes

import (
	pb "api-gateway/gen/product"
	"api-gateway/pkg/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductRoutes struct{
	ProductService *services.ProductService
}

func NewProductRoutes(service *services.ProductService) *ProductRoutes {
    return &ProductRoutes{ProductService: service}
}

func (pr *ProductRoutes)GetProductsRoute (c *gin.Context){
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
	res,err:=pr.ProductService.GetAllProductService(int32(page),int32(limit))
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
func (pr *ProductRoutes)GetProductByIdRoute(c *gin.Context){
	IdParam:=c.Param("id")
	Id,err:=strconv.Atoi(IdParam)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Bad Request",
		})
		return
	}
	res,err:=pr.ProductService.GetProductByIdService(Id)
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

func  (pr *ProductRoutes)DeleteProductRoute(c *gin.Context){
	idParam:=c.Param("id")
	Id,err:=strconv.Atoi(idParam)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"message":"Bad Request",
		})
		return
	}
	_,err=pr.ProductService.DeleteProductService(Id)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"Product Deleted successfully",
	})
}

func (pr *ProductRoutes)CreateProductRoute(c *gin.Context){
	 productReq:= &pb.Product{}
	if err:=c.ShouldBindJSON(productReq); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	res,err:=pr.ProductService.CreateProductService(productReq.Name,productReq.Description,productReq.Price)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated,gin.H{
		"message":"product created successfully",
		"data":res,
	})
}