package controller

import (
	"github.com/JevaPrahaysuma/managemnet.git/config"
	"github.com/JevaPrahaysuma/managemnet.git/models"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	users := []models.User{}
	if c.Param("id") == "" {
		config.DB.Find(&users)
	} else {
		config.DB.First(&users, c.Param("id"))
	}

	c.JSON(200, &users)
}
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	config.DB.Create(&user)
	c.JSON(200, &user)
}
func DeleteUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}
func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.JSON(200, &user)
}
