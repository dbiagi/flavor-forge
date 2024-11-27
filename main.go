package main

import (
	"context"
	"github.com/dbiagi/gororoba/src/config"
	"github.com/dbiagi/gororoba/src/handler"
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
	srv := createServer()
	configureGracefullShutdown(srv)
}

func createServer() *http.Server {
	router := mux.NewRouter()
	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      router,
		IdleTimeout:  time.Second * idleTimeout,
		ReadTimeout:  time.Second * readTimeout,
		WriteTimeout: time.Second * writeTimeout,
	}

	registerRoutesAndServe(router)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			slog.Error("Error starting server: %v\n", err)
		}
	}()

	return srv
}

func registerRoutesAndServe(router *mux.Router) {
	router.Use(config.TraceIdMiddleware)
	router.HandleFunc("/recipes", handler.GetRecipes).Methods("GET")
	router.HandleFunc("/health", handler.HealthCheck).Methods("GET")
	router.HandleFunc("/health/complete", handler.HealthCheckComplete).Methods("GET")
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
