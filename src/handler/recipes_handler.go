package handler

import (
	"time"

	"github.com/dbiagi/gororoba/src/converter"
	"github.com/dbiagi/gororoba/src/domain"
	"github.com/dbiagi/gororoba/src/repository"
	"github.com/google/uuid"
)

type RecipesHandlerInterface interface {
	GetRecipesByCategory(category string) []domain.Recipe
	CreateRecipe(r *domain.Recipe) *domain.Recipe
}

type RecipesHandler struct {
	RecipeRepository repository.RecipeRepositoryInterface
}

func NewRecipesHandler(r repository.RecipeRepositoryInterface) RecipesHandler {
	return RecipesHandler{RecipeRepository: r}
}

func (h RecipesHandler) GetRecipesByCategory(category string) []domain.Recipe {
	return h.RecipeRepository.GetRecipesByCategory(category)
}

func (h RecipesHandler) CreateRecipe(r *domain.Recipe) *domain.Recipe {
	r.Id = uuid.New().String()
	r.CreatedAt = time.Now()
	r.UpdatedAt = time.Now()
	m := converter.ToRecipeModel(*r)
	h.RecipeRepository.CreateRecipe(m)

	return r
}
