package usecase

import (
	"PruebaGo/internal/auth"
	user "PruebaGo/internal/domain/users"
	"PruebaGo/internal/domain/users/entities"
	"PruebaGo/internal/domain/users/http/models"
	validation "PruebaGo/pkg"
	"PruebaGo/pkg/password"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// usecase implementa la lógica de negocio para usuarios.
type usecase struct {
	repo user.Repository
}

// NewUsecase crea una nueva instancia de usecase.
func NewUsecase(repo user.Repository) user.Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) RegisterUser(ctx context.Context, request models.RegisterUserRequest) (models.RegisterResponse, error) {

	err := validation.ValidateUser(request.Username, request.Phone, request.Email, request.Password)
	if err != nil {
		return models.RegisterResponse{}, fmt.Errorf("validation failed: %w", err)
	}

	hashedPassword, err := password.HashPassword(request.Password)
	if err != nil {
		return models.RegisterResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}

	user := entities.User{
		ID:       uuid.New(),
		Username: request.Username,
		Phone:    request.Phone,
		Email:    request.Email,
		Password: hashedPassword,
	}

	_, err = u.repo.RegisterUser(ctx, user)
	if err != nil {
		return models.RegisterResponse{}, fmt.Errorf("failed to register user: %w", err)
	}

	// Crear y devolver la respuesta del registro
	response := models.RegisterResponse{
		Username: user.Username,
	}

	return response, nil
}

func (u *usecase) Login(ctx context.Context, request models.LoginRequest) (models.LoginResponse, error) {

	// Validar los datos de login
	err := validation.ValidateLoginData(request.Identifier, request.Password)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("validación fallida: %w", err)
	}

	// Buscar el usuario por identificador (username o email)
	user, err := u.repo.GetUserByIdentifier(ctx, request.Identifier)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("usuario no encontrado: %w", err)
	}

	// Verificar la contraseña
	if !password.CheckPasswordHash(request.Password, user.Password) {
		return models.LoginResponse{}, errors.New("contraseña incorrecta")
	}

	token, err := auth.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return models.LoginResponse{}, fmt.Errorf("error al generar el token: %w", err)
	}

	// Crear y devolver la respuesta del login
	return models.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Token:    token,
	}, nil
}
