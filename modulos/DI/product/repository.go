package product

import "database/sql"

type ProductRepositoryInterface interface {
	GetByID(id int) (Product, error)
}

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetByID(id int) (Product, error) {
	return Product{ID: id, Name: "Sample Product"}, nil
}
