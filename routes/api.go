package routes
import (
	"com.qmove/controllers/services"
	"github.com/gin-gonic/gin"
)
func InitApiRoutes(r *gin.Engine){
	r.GET("/ping", services.Index)
}
