package main

import (
	"log"
	"owlint/routes"
	"owlint/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDB()

	router := gin.Default()

	routes.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("error while starting server:", err)
	}
}
