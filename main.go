package main

import (
	routes "helpme-backend/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	router := gin.New()

	router.Use(CORSMiddleware())

	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.InterestRoutes(router)
	routes.HelpRoutes(router)
	routes.UserRoutes(router)
	routes.WishListRouter(router)

	port := ":" + os.Getenv("PORT")
	router.Run(port)
}
