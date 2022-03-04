package routes

import (

	controller "alef_education_devops_challenge/controllers"

	"github.com/gin-gonic/gin"
)

//TaskRouter function
func TaskRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/api/v1/tasks", controller.AddTask())
	// incomingRoutes.GET("/api/v1/helpee/wishlist/:user_id/:whishListId", controller.GetWishList())
	// incomingRoutes.GET("/api/v1/helpee/wishlist/:user_id", controller.GetAllWishListByUserID())
	// incomingRoutes.GET("/api/v1/helpee/count/wishlist/:user_id/:whishListId", controller.CountWhishList())
	// incomingRoutes.DELETE("/api/v1/helpee/wishlist/:user_id/:whishListId", controller.DeleteWishListByID())
	// incomingRoutes.PATCH("/api/v1/helpee/wishlist/:user_id/:whishListId", controller.UpdateWishList())
}
