package models

type User struct {
	Default
	Name           string `json:"name" gorm:"not null;"`
	Email          string `json:"email" gorm:"not null;unique"`
	Responsability string `json:"responsability"`
	Password       []byte `json:"-" gorm:"not null;"`
}
