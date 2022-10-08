package routes

import (
	"github.com/JevaPrahaysuma/managemnet.git/controller"
	"github.com/gin-gonic/gin"
)

func MenuRoute(router *gin.Engine) {
	router.GET("/", controller.GetMenu)
}
