package database

import (
	"database/sql"
)

type Category struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(id, name, description string) (Category, error) {
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)", id, name, description)
	if err != nil {
		return Category{}, err
	}

	category := Category{
		ID:          id,
		Name:        name,
		Description: description,
	}

	return category, nil
}

func (c *Category) GetAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]Category, 0)

	for rows.Next() {
		var id, name, description string
		err := rows.Scan(&id, &name, &description)
		if err != nil {
			return nil, err
		}
		category := Category{
			ID:          id,
			Name:        name,
			Description: description,
		}
		categories = append(categories, category)
	}

	return categories, nil
}
