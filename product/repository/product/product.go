package product

import (
	"database/sql"
	"errors"

	"github.com/larien/product/product/drivers/database"
	"github.com/larien/product/product/entity"
)

// Product represents the methods available in this repository layer
type Product interface {
	// Create inserts a product into the database
	Create(product *entity.Product) error
}

type repository struct {
	DB *database.DB
}

// New creates a new instance of Product repository to manipulate the database
func New(db *database.DB) Product {
	return &repository{db}
}


func (r *repository) Create(product *entity.Product) error {
	_, err := r.DB.NamedExec(`
		INSERT INTO products(id, price_in_cents, title, description)
		VALUES (:id, :price_in_cents, :title, :description)`, product)
	return err
}
