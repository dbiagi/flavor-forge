package tests

import (
	"gororoba/internal/controller"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RecipesControllerTestSuite struct {
	suite.Suite
	controller.RecipesController
}

func (s *HttpTestSuite) TestGetRecipesByCategory() {
	r, err := http.Get(s.BaseURI + "/recipes/by-category?category=snack")

	assert.Nil(s.T(), err)

	assert.Equal(s.T(), 200, r.StatusCode)
}
