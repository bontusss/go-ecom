package main

import (
	"log"
	"os"

	"github.com/bontusss/go-ecom/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// WARNING: PASSONG A POINTER MIGHT BE WRONG
	routes.UserRoutes(*router)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Server could not start", err.Error())
	}
}
