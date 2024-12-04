package main

import (
	"context"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/dbiagi/gororoba/src/config"
	"github.com/dbiagi/gororoba/src/controller"
	"github.com/dbiagi/gororoba/src/handler"
	"github.com/dbiagi/gororoba/src/repository"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

type Controllers struct {
	controller.HealthCheckController
	controller.RecipesController
}

func main() {
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development"
	}

	appConfig := config.LoadConfig(appEnv)

	config.ConfigureLogger(appConfig.AppConfig)

	slog.Info("Connecting to dynamoDB ....")
	dynamoDB := connectToDynamoDB(appConfig.AWSConfig)

	slog.Info("Starting server ....")
	srv, router := createServer(appConfig.WebConfig)

	slog.Info("Creating resources ....")
	appResources := createControllers(dynamoDB)

	slog.Info("Registering routes and serving ....")
	registerRoutesAndServe(router, appResources)

	slog.Info("Configuring graceful shutdown ....")
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
		Addr:         ":" + os.Getenv("PORT"),
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

func createControllers(db *dynamodb.DynamoDB) Controllers {
	recipeRepository := repository.NewRecipeRepository(db)
	healthCheckHandler := handler.NewHealthCheckHandler()
	recipeHandler := handler.NewRecipesHandler(recipeRepository)

	return Controllers{
		RecipesController:     controller.NewRecipesController(recipeHandler),
		HealthCheckController: controller.NewHealthCheckController(healthCheckHandler),
	}
}

func registerRoutesAndServe(router *mux.Router, controllers Controllers) {
	router.Use(config.TraceIdMiddleware)
	router.HandleFunc("/health", controller.HandleRequest(controllers.HealthCheckController.Check)).Methods("GET")
	router.HandleFunc("/health/complete", controller.HandleRequest(controllers.HealthCheckController.CheckComplete)).Methods("GET")
	router.HandleFunc("/recipes", controller.HandleRequest(controllers.RecipesController.GetRecipes)).Methods("GET")
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
