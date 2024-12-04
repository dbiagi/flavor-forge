package domain

import (
	"time"
)

type RecipeCategory string

const (
	Dessert    RecipeCategory = "Dessert"
	Soup       RecipeCategory = "Soup"
	MainCourse RecipeCategory = "Main Course"
	SideDish   RecipeCategory = "Side Dish"
	Drink      RecipeCategory = "Drink"
	Appetizer  RecipeCategory = "Appetizer"
	Salad      RecipeCategory = "Salad"
)

// PK = Category, SK = Id#UpdatedAt
// GSI = Id
// GSI = Slug
type Recipe struct {
	Id             string       `json:"id"`
	Title          string       `json:"title"`
	Description    string       `json:"description"`
	Servings       int          `json:"servings"`
	PrepTime       int          `json:"prepTime"`
	Slug           string       `json:"slug"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
	Ingredients    []Ingredient `json:"ingredients"`
	Category       string       `json:"category"`
	IdAndUpdatedAt string       `json:"id#updatedAt"`
}

type Ingredient struct {
	Name        string
	Quantity    string
	MeasureUnit MeasureUnit
}
