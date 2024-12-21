package repository_mock

import (
	"github.com/dbiagi/gororoba/src/domain"
	"github.com/dbiagi/gororoba/src/model"
)

type MockRecipeRepository struct {
	Calls int
}

func (m *MockRecipeRepository) GetRecipesByCategory(category string) []domain.Recipe {
	m.Calls++
	return []domain.Recipe{}
}

func (r *MockRecipeRepository) CreateRecipe(recipe model.RecipeModel) *domain.Error {
	return nil
}
