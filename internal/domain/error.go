package domain

type Error struct {
	Message string
	Cause   error
}
