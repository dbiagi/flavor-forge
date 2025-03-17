package internal

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"gororoba/internal/config"
	"gororoba/internal/controller"
	"gororoba/internal/handler"
	"gororoba/internal/repository"
	"gororoba/internal/utils"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	signalsToListenTo = []os.Signal{
		syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM,
	}
)

type controllers struct {
	controller.HealthCheckController
	controller.RecipesController
}

type AppServer struct {
	config.Configuration
	dynamodb.DynamoDB
	*http.Server
	*mux.Router
}

func NewAppServer(appConfig config.Configuration) AppServer {
	return AppServer{
		Configuration: appConfig,
	}
}

func (s *AppServer) Start() {
	startTime := time.Now()
	config.ConfigureLogger(s.Configuration.AppConfig)

	slog.Info("Connecting to dynamoDB ....")
	dynamoDB := connectToDynamoDB(s.Configuration.AWSConfig)
	s.DynamoDB = *dynamoDB

	slog.Info("Starting server ....")
	srv, router := createServer(s.Configuration.WebConfig)
	s.Server = srv
	s.Router = router

	slog.Info("Creating resources ....")
	appResources := createControllers(dynamoDB)

	slog.Info("Registering routes and serving ....")
	registerRoutesAndMiddlewares(router, appResources)

	slog.Info(fmt.Sprintf("Application ready. Time elapsed: %v", time.Since(startTime)))

	if !s.Configuration.WebConfig.GracefulShutdownDisabled {
		configureGracefullShutdown(srv, s.Configuration.WebConfig)
	}
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
		Addr:         fmt.Sprintf(":%d", webConfig.Port),
		Handler:      router,
		IdleTimeout:  webConfig.IdleTimeout,
		ReadTimeout:  webConfig.ReadTimeout,
		WriteTimeout: webConfig.WriteTimeout,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err.Error() != "http: Server closed" {
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

func registerRoutesAndMiddlewares(router *mux.Router, controllers controllers) {
	router.Use(config.TraceIdMiddleware)
	router.Use(mux.CORSMethodMiddleware(router))
	router.HandleFunc("/health", utils.HandleRequest(controllers.HealthCheckController.Check)).Methods("GET")
	router.HandleFunc("/health/complete", utils.HandleRequest(controllers.HealthCheckController.CheckComplete)).Methods("GET")
	router.HandleFunc("/recipes/by-category", utils.HandleRequest(controllers.RecipesController.GetRecipesByCategory)).Methods("GET")
	router.HandleFunc("/recipes/suggestion", utils.HandleRequest(controllers.RecipesController.GetSuggestion)).Methods("GET")
}

func configureGracefullShutdown(server *http.Server, webConfig config.WebConfig) {
	slog.Info("Configuring graceful shutdown.")

	c := make(chan os.Signal, 1)
	signal.Notify(c, signalsToListenTo...)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), webConfig.ShutdownTimeout)
	defer cancel()

	server.Shutdown(ctx)
	slog.Info("Shutting down server")
	os.Exit(0)
}

func (s *AppServer) ForceShutdown() {
	s.Server.Shutdown(context.Background())
}
