package model

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Price       float64   `json:"price"`
	Category    *Category `json:"category"`
}
