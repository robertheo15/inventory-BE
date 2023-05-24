package inventory

import (
	"gorm.io/gorm"
)

type PostgresInventoryRepository struct {
	db *gorm.DB
}

func NewPostgresInventoryRepository(db *gorm.DB) *PostgresInventoryRepository {
	return &PostgresInventoryRepository{
		db: db,
	}
}
