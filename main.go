package main

import (
	"QUZHIYOU/cache"
	"QUZHIYOU/models"
	"QUZHIYOU/routers"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	cache.Redis()
	models.Initialized()
}

func main() {
	//gin.SetMode(gin.DebugMode)
	defer models.CloseDb()
	router := routers.InitRouter()
	router.Run(":8080")
}
