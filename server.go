package main

import (
	"com.qmove/database"
	"com.qmove/routes"
	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	routes.InitRoutes(r)
	database.InitDatabaseConnection()
	r.Run()
}