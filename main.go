package main

import (
	"fmt"
	"os"

	"github.com/dbiagi/gororoba/src/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Loading environment variables")
	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file: %v\n", envErr)
		return
	}

	fmt.Println("Var: ", os.Getenv("POSTGRES_URL"))

	fmt.Println("Starting server ....")
	router := gin.Default()
	registerHandlers(router)

	routerErr := router.Run(":" + os.Getenv("PORT"))
	if routerErr != nil {
		fmt.Printf("Error starting server: %v\n", routerErr)
	}
}

func registerHandlers(router *gin.Engine) {
	router.GET("/recipes", handler.GetRecipes)
	router.GET("/health", handler.HealthCheck)
	router.GET("/health/complete", handler.HealthCheckComplete)
}
