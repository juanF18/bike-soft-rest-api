package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/juanF18/bike-soft-rest-api.git/routes"
)

func main() {
	router := gin.Default()

	routes.Routes(router)
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
