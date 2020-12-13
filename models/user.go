package models

type User struct {
	Login      string `json:"login" db:"login"`
	FirstName  string `json:"firstName" db:"first_name"`
	SecondName string `json:"secondName" db:"second_name"`
	Email      string `json:"email" db:"email"`
}
