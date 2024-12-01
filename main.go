package main

import (
	"context"
	"github.com/dbiagi/gororoba/src/config"
	"github.com/dbiagi/gororoba/src/controller"
	"github.com/dbiagi/gororoba/src/handler"
	"github.com/dbiagi/gororoba/src/repository"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const (
	idleTimeout  = 10
	writeTimeout = 10
	readTimeout  = 10
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		slog.Info("Error loading .env file: %v\n", envErr)
		return
	}

	appEnv := os.Getenv("APP_ENV")
	config.ConfigureLogger(appEnv)

	slog.Info("Starting server ....")
	srv, router := createServer()

	slog.Info("Creating resources ....")
	appResources := createControllers()

	slog.Info("Registering routes and serving ....")
	registerRoutesAndServe(router, appResources)

	slog.Info("Configuring graceful shutdown ....")
	configureGracefullShutdown(srv)
}

func createServer() (*http.Server, *mux.Router) {
	router := mux.NewRouter()
	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      router,
		IdleTimeout:  time.Second * idleTimeout,
		ReadTimeout:  time.Second * readTimeout,
		WriteTimeout: time.Second * writeTimeout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			slog.Error("Error starting server: %v\n", err)
		}
	}()

	return srv, router
}

type Controllers struct {
	controller.HealthCheckController
	controller.RecipesController
}

func createControllers() Controllers {
	db, err := config.ConnectToDatabase()
	if err != nil {
		slog.Error("Error connecting to database: %v\n", err)
	}

	recipeRepository := repository.NewRecipeRepository(*db)
	healthCheckHandler := handler.NewHealthCheckHandler(*db)
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

func configureGracefullShutdown(server *http.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*writeTimeout)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("Shutting down server")
	os.Exit(0)
}
