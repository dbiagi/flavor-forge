package domain

import (
	"time"
)

type Recipe struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
}
