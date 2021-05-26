package routes

import (
	"github.com/gin-gonic/gin"
	)
func InitRoutes(r *gin.Engine) {
	InitApiRoutes(r)
	InitWebRoutes(r)
}
