package main

import (
	routes "alef_education_devops_challenge/routes"
	"log"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.Use(gin.Logger())
	routes.TaskRouter(router)

	es, err := elasticsearch.NewDefaultClient()
	// log.Println(elasticsearch.Version)
	// log.Println(es.Info())

	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)

	port := ":" + os.Getenv("PORT")
	router.Run(port)
}
