package app

import (
	v1 "PruebaGo/internal/domain/users/http"
	postgres "PruebaGo/internal/domain/users/infraestructure"
	"PruebaGo/internal/domain/users/repository"
	usecase "PruebaGo/internal/domain/users/usecases"
	"context"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

// UserService se encarga de inicializar el servicio de usuarios.
// Esta función inicializa las dependencias necesarias para el servicio de usuarios,
// tales como el almacenamiento, el repositorio, los casos de uso, y las rutas HTTP.
func UserService(ctx context.Context, db *pgxpool.Pool, router *http.ServeMux) error {
	storage := postgres.NewPgxStorage(db)

	repo := repository.NewUserRepository(storage)

	uc := usecase.NewUsecase(repo)

	h := v1.NewHandler(uc)

	h.UserRoutes(router)

	return nil
}
