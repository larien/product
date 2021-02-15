package product

import (
	"database/sql"
	"errors"

	"github.com/larien/product/product/drivers/database"
	"github.com/larien/product/product/entity"
)

// Product represents the methods available in this repository layer
type Product interface {
	// Get obtains a product by its ID from the database
	Get(id string) (*entity.Product, error)
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

func (r *repository) Get(id string) (*entity.Product, error) {
	var product entity.Product
	err := r.DB.Get(&product, `
		SELECT id, price_in_cents, title, description, created_at, updated_at, deleted_at
		FROM products
		WHERE id = $1
		LIMIT 1`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &product, nil
}


func (r *repository) Create(product *entity.Product) error {
	_, err := r.DB.NamedExec(`
		INSERT INTO products(id, price_in_cents, title, description)
		VALUES (:id, :price_in_cents, :title, :description)`, product)
	return err
}
