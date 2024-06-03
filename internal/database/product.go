package database

import "database/sql"

type Product struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
	Price       float64
	CategoryID  string
}

func NewProduct(db *sql.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Create(id, name, description, categoryID string, price float64) (Product, error) {
	_, err := p.db.Exec("INSERT INTO products (id, name, description, category_id, price) VALUES ($1, $2, $3, $4, $5)", id, name, description, categoryID, price)
	if err != nil {
		return Product{}, err
	}

	product := Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		CategoryID:  categoryID,
	}

	return product, nil
}

func (p *Product) GetAll() ([]Product, error) {
	rows, err := p.db.Query("SELECT id, name, description, price, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]Product, 0)

	for rows.Next() {
		var id, name, description, categoryID string
		var price float64
		err := rows.Scan(&id, &name, &description, &price, &categoryID)
		if err != nil {
			return nil, err
		}
		product := Product{
			ID:          id,
			Name:        name,
			Description: description,
			Price:       price,
			CategoryID:  categoryID,
		}
		products = append(products, product)
	}

	return products, nil
}
