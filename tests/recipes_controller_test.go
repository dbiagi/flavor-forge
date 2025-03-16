package tests

import (
	"gororoba/internal/controller"

	"github.com/stretchr/testify/suite"
)

type RecipesControllerTestSuite struct {
	suite.Suite
	controller.RecipesController
}

func (s *RecipesControllerTestSuite) SetupSuite() {

	// r := repository.NewRecipeRepository()
	// h := handler.NewRecipesHandler()
	// c := controller.NewRecipesController()
	// ts := RecipesControllerTestSuite{}
}

func (s *HttpTestSuite) TestGetRecipesByCategory() {

}

func NewHandlerFuncWithMiddleware() {

}
