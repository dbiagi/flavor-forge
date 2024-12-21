package controller

import (
	"github.com/dbiagi/gororoba/src/handler"
	"net/http"
)

type RecipesController struct {
	RecipesHandler handler.RecipesHandlerInterface
}

func NewRecipesController(h handler.RecipesHandlerInterface) RecipesController {
	return RecipesController{RecipesHandler: h}
}

func (rc *RecipesController) GetRecipes(w http.ResponseWriter, r *http.Request) HttpResponse {
	category := r.URL.Query().Get("category")
	recipes := rc.RecipesHandler.GetRecipesByCategory(category)
	return HttpResponse{Body: recipes}
}
