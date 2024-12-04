package controller

import (
	"github.com/dbiagi/gororoba/src/handler"
	"net/http"
)

type RecipesController struct {
	handler.RecipesHandler
}

func NewRecipesController(h handler.RecipesHandler) RecipesController {
	return RecipesController{RecipesHandler: h}
}

func (rc *RecipesController) GetRecipes(w http.ResponseWriter, r *http.Request) HttpResponse {
	category := r.URL.Query().Get("category")
	recipes := rc.RecipesHandler.GetRecipesByCategory(category)
	return HttpResponse{Body: recipes}
}
