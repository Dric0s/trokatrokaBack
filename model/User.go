package model

import "github.com/google/uuid"

type UserData struct {
	IdUser      uuid.UUID `json:"idUser"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	Phone       string    `json:"phone"`
	TotalRating float32   `json:"totalRating"`
	Books       []Book    `json:"books"`
}
