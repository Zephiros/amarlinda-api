package models

type OrderOut struct {
  Order
  ClientId int `json:"client_id" gorm:"not null"`
  Client   Client `gorm:"not null"`
}
