package model

import "time"

type Category struct {
	IdCategory  time.Time `json:"idCategory"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Books       []Book    `json:"books"`
}
