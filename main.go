package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gororoba/config"
	"gororoba/controller"
	"gororoba/handler"
	"gororoba/repository"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gorilla/mux"
)

type controllers struct {
	controller.HealthCheckController
	controller.RecipesController
}

func main() {
	startTime := time.Now()

	env := os.Args[1]
	if env == "" {
		env = config.DevelopmentEnv
	}

	appConfig := config.LoadConfig(env)

	config.ConfigureLogger(appConfig.AppConfig)

	slog.Info("Connecting to dynamoDB ....")
	dynamoDB := connectToDynamoDB(appConfig.AWSConfig)

	slog.Info("Starting server ....")
	srv, router := createServer(appConfig.WebConfig)

	slog.Info("Creating resources ....")
	appResources := createControllers(dynamoDB)

	slog.Info("Registering routes and serving ....")
	registerRoutesAndServe(router, appResources)

	slog.Info(fmt.Sprintf("Application ready. Time elapsed: %v", time.Since(startTime)))

	slog.Info("Configuring graceful shutdown.")
	configureGracefullShutdown(srv, appConfig.WebConfig)
}

func connectToDynamoDB(awsConfig config.AWSConfig) *dynamodb.DynamoDB {
	dynamoDB, err := config.CreateDynamoDBConnection(awsConfig)
	if err != nil {
		slog.Error("Error connecting to dynamodb.", slog.String("error", err.Message))
		panic(err)
	}

	return dynamoDB
}

func createServer(webConfig config.WebConfig) (*http.Server, *mux.Router) {
	router := mux.NewRouter()
	srv := &http.Server{
		Addr:         ":" + webConfig.Port,
		Handler:      router,
		IdleTimeout:  webConfig.IdleTimeout,
		ReadTimeout:  webConfig.ReadTimeout,
		WriteTimeout: webConfig.WriteTimeout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			slog.Error("Error starting server.", slog.String("error", err.Error()))
		}
	}()

	return srv, router
}

func createControllers(db *dynamodb.DynamoDB) controllers {
	recipeRepository := repository.NewRecipeRepository(db)
	healthCheckHandler := handler.NewHealthCheckHandler()
	suggestionHandler := handler.NewSuggestionHandler()
	recipeHandler := handler.NewRecipesHandler(recipeRepository, suggestionHandler)

	return controllers{
		RecipesController:     controller.NewRecipesController(recipeHandler),
		HealthCheckController: controller.NewHealthCheckController(healthCheckHandler),
	}
}

func registerRoutesAndServe(router *mux.Router, controllers controllers) {
	router.Use(config.TraceIdMiddleware)
	router.HandleFunc("/health", controller.HandleRequest(controllers.HealthCheckController.Check)).Methods("GET")
	router.HandleFunc("/health/complete", controller.HandleRequest(controllers.HealthCheckController.CheckComplete)).Methods("GET")
	router.HandleFunc("/recipes/by-category", controller.HandleRequest(controllers.RecipesController.GetRecipesByCategory)).Methods("GET")
	router.HandleFunc("/recipes/suggestion", controller.HandleRequest(controllers.RecipesController.GetSuggestion)).Methods("GET")
}

func configureGracefullShutdown(server *http.Server, webConfig config.WebConfig) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), webConfig.ShutdownTimeout)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("Shutting down server")
	os.Exit(0)
}
