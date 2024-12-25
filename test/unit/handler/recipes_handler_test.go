package handler_test

import (
	"testing"

	"github.com/dbiagi/gororoba/src/handler"
	mock_repository "github.com/dbiagi/gororoba/test/mocks"
	fixtures_test "github.com/dbiagi/gororoba/test/unit/fixtures"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testSetup struct {
	recipeHandler        *handler.RecipesHandler
	recipeRepositoryMock *mock_repository.MockRecipeRepositoryInterface
}

func setup(t *testing.T) testSetup {
	ctrl := gomock.NewController(t)
	rr := mock_repository.NewMockRecipeRepositoryInterface(ctrl)
	h := handler.NewRecipesHandler(rr)

	return testSetup{
		recipeHandler:        &h,
		recipeRepositoryMock: rr,
	}
}

func TestGetRecipesByCategory(t *testing.T) {
	t.Run("Should return a list of recipes by category", func(t *testing.T) {
		// Given
		s := setup(t)
		c := "Dessert"
		r := fixtures_test.GetRecipesWithCategory(c)
		s.recipeRepositoryMock.EXPECT().GetRecipesByCategory(c).Return(r)

		// When
		result := s.recipeHandler.GetRecipesByCategory(c)

		// Then
		assert.GreaterOrEqual(t, len(result), 1)
	})
}
