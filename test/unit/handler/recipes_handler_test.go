package handler_test

import (
	"testing"

	"github.com/dbiagi/gororoba/src/domain"
)

type MockRecipeRepository interface {
	GetRecipesByCategory(category string) []domain.Recipe
	CreateRecipe(recipe interface{}) *domain.Error
}

func TestGetRecipesByCategory(t *testing.T) {
	// t.Run("Should return a list of recipes by category", func(t *testing.T) {

	// 	h := handler.NewRecipesHandler(interface{})

	// 	category := "Dessert"
	// 	expectedRecipes := []domain.Recipe{}

	// 	recipes := h.GetRecipesByCategory(category)

	// 	assert.Equal(t, expectedRecipes, recipes)
	// })
}
