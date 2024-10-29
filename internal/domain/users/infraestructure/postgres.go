// Paquete encargado de las conexiones a postgres
package postgres

import "github.com/jackc/pgx/v5/pgxpool"

// PgxStorage gestiona el pool de conexiones a PostgreSQL.
type PgxStorage struct {
	DbPool *pgxpool.Pool
}

// NewPgxStorage devuelve una nueva instancia de PgxStorage.
func NewPgxStorage(dbPool *pgxpool.Pool) *PgxStorage {
	return &PgxStorage{DbPool: dbPool}
}
