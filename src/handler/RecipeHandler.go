package handler

import (
	"encoding/json"
	"github.com/dbiagi/gororoba/src/domain"
	"log/slog"
	"net/http"
	"time"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	slog.InfoContext(ctx, "Getting recipes")
	err := json.NewEncoder(w).Encode(Recipes())
	if err != nil {
		slog.Error("Error encoding recipes: %v\n", err)
		return
	}
}

func Recipes() []domain.Recipe {
	return []domain.Recipe{
		{
			Id:          1,
			Title:       "Bolo de cenoura",
			Description: "Bolo de cenoura com cobertura de chocolate",
			CreatedAt:   time.Now(),
		},
		{
			Id:          2,
			Title:       "Bolo de chocolate",
			Description: "Bolo de chocolate com cobertura de brigadeiro",
			CreatedAt:   time.Now(),
		},
	}
}
