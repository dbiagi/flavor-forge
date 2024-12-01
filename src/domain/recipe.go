package domain

import (
	"github.com/google/uuid"
	"time"
)

type Recipe struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Servings    int       `json:"servings"`
	PrepTime    int       `json:"prepTime"`
	Slug        string    `json:"slug"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
