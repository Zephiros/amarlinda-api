package models

type OrderItem struct {
	Default
	OrderId       int     `json:"order_id" gorm:"not null"`
	Order         Order   `gorm:"not null"`
	ProductId     int     `json:"product_id" gorm:"not null"`
	Product       Product `gorm:"not null"`
	SubTotalValue float32 `json:"sub_total_value" gorm:"type:decimal(10,2);not null;default:0.0"`
	TotalValue    float32 `json:"total_value" gorm:"type:decimal(10,2);not null;default:0.0"`
	Observation   string  `json:"observation"`
}
