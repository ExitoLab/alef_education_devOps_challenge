package main

import (
	routes "alef_education_devops_challenge/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())

	routes.TaskRouter(router)

	port := ":" + os.Getenv("PORT")
	router.Run(port)
}
