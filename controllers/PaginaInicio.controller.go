package controllers

import "github.com/gin-gonic/gin"

func HandleHomepage(ctx *gin.Context) {
	ctx.String(200, "Hola mundo")
}
