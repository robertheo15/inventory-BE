package repository

import (
	"gorm.io/gorm"
)

type PosgtresRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PosgtresRepository {
	return &PosgtresRepository{
		db: db,
	}
}
