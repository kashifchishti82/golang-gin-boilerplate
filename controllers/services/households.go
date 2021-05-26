package services

import (
	"com.qmove/database"
	"com.qmove/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context){
	var households [] models.HouseHold
	database.DB.Preload("Contacts").Find(&households)
	c.JSON(http.StatusOK, gin.H{"data": households})
}
