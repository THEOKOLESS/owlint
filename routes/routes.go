package routes

import (
	"owlint/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/target/:targetId/comments", controllers.AddComment)
	router.GET("/target/:targetId/comments", controllers.GetComments)
}
