package model

type Category struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type CategoryInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}
