package main

import (
	"QUZHIYOU/models"
	"QUZHIYOU/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	models.Initialized()
}

func main() {
	gin.SetMode(gin.DebugMode)
	defer models.CloseDb()
	router := routers.InitRouter()
	router.Run(":8080")
}
