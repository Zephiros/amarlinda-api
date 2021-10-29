package seeds

import (
	"github.com/Zephiros/amarlinda-api/pkg/seed"
	"gorm.io/gorm"
)

func All() []seed.Seed {
	return []seed.Seed{{
		Name: "CreateJane",
		Run: func(db *gorm.DB) error {
			return CreatePayment(db, "Boleto Bancário")
		},
	}, {
		Name: "CreateJohn",
		Run: func(db *gorm.DB) error {
			return CreatePayment(db, "Dinheiro")
		},
	}}
}
