package repository

import (
	"github.com/dbiagi/gororoba/src/config"
	"github.com/dbiagi/gororoba/src/domain"
	"log/slog"
)

type RecipeRepository struct {
	config.Database
}

func NewRecipeRepository(db config.Database) RecipeRepository {
	return RecipeRepository{Database: db}
}

func (r *RecipeRepository) GetRecipes() []domain.Recipe {
	rows, err := r.Database.Query("SELECT id, title, description, servings, prep_time, slug, created_at, updated_at FROM recipe LIMIT 10")

	if err != nil {
		slog.Error("Error querying recipes: %v\n", err)
		return []domain.Recipe{}
	}

	var recipes []domain.Recipe
	var recipe domain.Recipe
	for rows.Next() {
		err := rows.Scan(
			&recipe.Id,
			&recipe.Title,
			&recipe.Description,
			&recipe.Servings,
			&recipe.PrepTime,
			&recipe.Slug,
			&recipe.CreatedAt,
			&recipe.UpdatedAt)
		if err != nil {
			slog.Error("Error scanning recipe: %v\n", err)
			return []domain.Recipe{}
		}
		recipes = append(recipes, recipe)
	}

	defer rows.Close()

	return recipes
}
