package routes

import (
	"github.com/bontusss/go-ecom/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes gin.Engine) {
	routes.POST("/user/register", controllers.SignUp())
	routes.POST("/user/login", controllers.Login)
	routes.POST("/admin/addproduct", controllers.ProductViewerAdmin)
	routes.GET("/user/productview", controllers.Searchproduct)
	routes.GET("/users/search", controllers.SearchproductByQuery)
}
