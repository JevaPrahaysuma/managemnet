package main

import (
	"github.com/JevaPrahaysuma/managemnet.git/config"
	"github.com/JevaPrahaysuma/managemnet.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	//routes.UserRoute(router)
	routes.MenuRoute(router)
	router.Run(":8080")
}
