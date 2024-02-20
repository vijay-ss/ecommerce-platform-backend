package routes

import (
	"github.com/gin-gonic/gin"
	
	controller "github.com/vijay-ss/ecommerce-platform-backend/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controller.Signup())
	incomingRoutes.POST("/users/login", controller.Login())
	incomingRoutes.POST("/admin/addproduct", controller.ProductViewerAdmin())
	incomingRoutes.GET("/users/productview", controller.SearchProduct())
	incomingRoutes.GET("/users/search", controller.SearchProductByQuery())
}