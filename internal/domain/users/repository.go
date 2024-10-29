package user

import (
	"PruebaGo/internal/domain/users/entities"
	"context"
)

// Repository define las operaciones de persistencia de usuarios.
type Repository interface {
	Get
	Register
}

type Register interface {
	RegisterUser(ctx context.Context, user entities.User) (entities.User, error)
}

type Get interface {
	GetUserByIdentifier(ctx context.Context, identifier string) (entities.User, error)
}
