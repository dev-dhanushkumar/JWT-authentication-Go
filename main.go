package main

import (
	"fmt"
	"log"
	"os"
	"sam0307204/jwt-Authentication/controller"
	"sam0307204/jwt-Authentication/database"
	"sam0307204/jwt-Authentication/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadTheEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading in env file")
	}
}

func main() {
	LoadTheEnv()
	DB_Connect := os.Getenv("DB_Connect")
	fmt.Println("Server loading...")

	//Database Conection
	database.Connect(DB_Connect)
	database.Migrate()

	//Router connection
	router := initRouter()
	router.Run(":8080")
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controller.GenerateToken)
		api.POST("/user/register", controller.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controller.Ping)
		}
		return router
	}
}
