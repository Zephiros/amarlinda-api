package models

type OrderIn struct {
	Order
	Supplier string `json:"supplier"`
}
