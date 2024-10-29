package repository

import (
	user "PruebaGo/internal/domain/users"
	"PruebaGo/internal/domain/users/entities"
	postgres "PruebaGo/internal/domain/users/infraestructure"
	"context"
	"fmt"
	"log"
)

// userRepository implementa la interfaz user.Repository para interactuar con PostgreSQL.
type userRepository struct {
	storage *postgres.PgxStorage
}

// NewUserRepository crea una nueva instancia de userRepository.
func NewUserRepository(storage *postgres.PgxStorage) user.Repository {
	return &userRepository{storage: storage}
}

func (ur *userRepository) RegisterUser(ctx context.Context, ru entities.User) (entities.User, error) {
	query := "INSERT INTO users (id, username, phone, email, password) VALUES ($1, $2, $3, $4, $5) RETURNING username"

	log.Printf("Registrando usuario con ID: %s, Username: %s, Phone: %s, Email: %s, Password: [HASHED]", ru.ID, ru.Username, ru.Phone, ru.Email)

	var user entities.User

	err := ur.storage.DbPool.QueryRow(ctx, query, ru.ID, ru.Username, ru.Phone, ru.Email, ru.Password).Scan(&user.Username)
	if err != nil {
		return user, fmt.Errorf("failed to register user: %w", err)
	}
	return user, nil
}

func (ur *userRepository) GetUserByIdentifier(ctx context.Context, identifier string) (entities.User, error) {
	query := "SELECT username, email, password FROM users WHERE username = $1 OR email = $1"

	var user entities.User
	err := ur.storage.DbPool.QueryRow(ctx, query, identifier).Scan(&user.Username, &user.Email, &user.Password)
	if err != nil {
		return user, fmt.Errorf("failed to find user: %w", err)
	}

	return user, nil
}
