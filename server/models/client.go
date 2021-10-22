package models

type Client struct {
  Default
  Name  string `json:"name" gorm:"not null;"`
  Phone string `json:"phone" gorm:"type:varchar(20);unique"`
}
