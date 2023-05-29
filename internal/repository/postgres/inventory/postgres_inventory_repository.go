package inventory

import (
	"database/sql"
	"inventory-app-be/internal/repository/postgres/sqlc"
)

type PostgresInventoryRepository struct {
	db *sqlc.Queries
}

func NewPostgresInventoryRepository(db *sql.DB) *PostgresInventoryRepository {
	return &PostgresInventoryRepository{
		db: sqlc.New(db),
	}
}
