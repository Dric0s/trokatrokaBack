package model

import (
	"github.com/google/uuid"
	"time"
)

type Book struct {
	IdBook     uuid.UUID `json:"idBook"`
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Isbn       string    `json:"isbn"`
	Publisher  string    `json:"publisher"`
	Edition    string    `json:"edition"`
	Model      int32     `json:"model"`
	Language   string    `json:"language"`
	Reason     string    `json:"reason"`
	BuyPrice   float64   `json:"buyPrice"`
	BuyDate    time.Time `json:"buyDate"`
	User       UserData  `json:"user"`
	Categories Category  `json:"categories"`
}
