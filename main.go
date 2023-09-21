package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"go-example-api/app/api/router"
	"go-example-api/config"
	"os"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading .env file. Error: ", err)
	}
	config.InitLog()
}

func main() {
	Init()

	db := config.ConnectDB()
	appInitialization := config.InitAPI(db)

	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	router.Init(appInitialization, ginEngine)

	port := os.Getenv("PORT")
	err2 := ginEngine.Run(":" + port)
	if err2 != nil {
		log.Fatal("Error while setting up server. Error: ", err2)
	}
	log.Info("Server listening on PORT: ", port)
}
