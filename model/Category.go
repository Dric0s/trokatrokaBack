package model

type Category struct {
	IdCategory  uint64 `json:"idCategory"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Books       []Book `json:"books"`
}
