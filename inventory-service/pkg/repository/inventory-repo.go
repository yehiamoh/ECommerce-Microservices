package repository

import "database/sql"

type InventoryRepository struct {
	db *sql.DB
}
func NewInventoryRepository(db *sql.DB)*InventoryRepository{
	return &InventoryRepository{
		db: db,
	}
}
