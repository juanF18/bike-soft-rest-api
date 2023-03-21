package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/juanF18/bike-soft-rest-api.git/controllers"
)

func Routes(router *gin.Engine) {
	rustaInicio(router)
}

func rustaInicio(router *gin.Engine) {
	router.GET("/", controllers.HandleHomepage)
}
