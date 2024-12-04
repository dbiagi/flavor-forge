package handler

import (
	"github.com/dbiagi/gororoba/src/domain"
	"github.com/dbiagi/gororoba/src/repository"
)

type RecipesHandler struct {
	repository.RecipeRepository
}

func NewRecipesHandler(r repository.RecipeRepository) RecipesHandler {
	return RecipesHandler{RecipeRepository: r}
}

func (h *RecipesHandler) GetRecipesByCategory(category string) []domain.Recipe {
	return h.RecipeRepository.GetRecipesByCategory(category)
}
