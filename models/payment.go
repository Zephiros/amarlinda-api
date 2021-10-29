package models

type Payment struct {
	Default
	Name string `json:"name" gorm:"type:varchar(30)not null;unique"`
}
