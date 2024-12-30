package handler_test

import (
	"testing"

	"gororoba/handler"
	testdata_fixtures "gororoba/testdata/fixtures"
	testdata_mocks "gororoba/testdata/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type testSetup struct {
	recipeHandler        *handler.RecipesHandler
	recipeRepositoryMock *testdata_mocks.MockRecipeRepositoryInterface
}

func setup(t *testing.T) testSetup {
	ctrl := gomock.NewController(t)
	rr := testdata_mocks.NewMockRecipeRepositoryInterface(ctrl)
	h := handler.NewRecipesHandler(rr)

	return testSetup{
		recipeHandler:        &h,
		recipeRepositoryMock: rr,
	}
}

func TestGetRecipesByCategory(t *testing.T) {

	// Given
	s := setup(t)
	c := "Dessert"
	r := testdata_fixtures.GetRecipesWithCategory(c)
	s.recipeRepositoryMock.EXPECT().GetRecipesByCategory(c).Return(r)

	// When
	result := s.recipeHandler.GetRecipesByCategory(c)

	// Then
	assert.GreaterOrEqual(t, len(result), 1)

}
