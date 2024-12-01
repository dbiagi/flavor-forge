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

func (h *RecipesHandler) GetRecipes() []domain.Recipe {
	return h.RecipeRepository.GetRecipes()
}
