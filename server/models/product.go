package models

type Product struct {
  Default
  Code            *string `json:"code" gorm:"unique"`
  Name            *string `json:"name"`
  Description     string `json:"description"`
  AmountAvailable int64  `json:"amount_available" gorm:"not null;default:0"`
  Price           float32 `json:"price" gorm:"type:decimal(10,2);not null;default:0.0"`
  PricePromotion  float32 `json:"price_promotion" gorm:"type:decimal(10,2);"`
  Active          bool   `json:"active" gorm:"default:true"`
  Promotion       bool   `json:"promotion" gorm:"default:false"`
}
