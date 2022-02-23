package model

type UserData struct {
	IdUser      uint64  `json:"idUser"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Phone       string  `json:"phone"`
	TotalRating float32 `json:"totalRating"`
	Books       []Book  `json:"books"`
}
