package resources

import (
	"github.com/dbiagi/gororoba/src/config"
	"github.com/dbiagi/gororoba/src/controller"
	"github.com/dbiagi/gororoba/src/handler"
	"github.com/dbiagi/gororoba/src/repository"
)

type Resources struct {
	config.Database
	handler.HealthCheckHandler
	handler.RecipesHandler
	controller.HealthCheckController
	controller.RecipesController
	repository.RecipeRepository
}
