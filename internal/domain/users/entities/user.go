package entities

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	Username string
	Phone    string
	Email    string
	Password string
}
