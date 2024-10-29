package user

import (
	"PruebaGo/internal/domain/users/http/models"
	"context"
)

// Usecase define la lógica de negocio para el manejo de usuarios.
type Usecase interface {
	RegisterUser(ctx context.Context, request models.RegisterUserRequest) (models.RegisterResponse, error)

	Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error)
}
