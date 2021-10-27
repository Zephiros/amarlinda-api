package seeds

import (
		"github.com/Zephiros/amarlinda/pkg/seed"
		"gorm.io/gorm"
)

func All() []seed.Seed {
		return []seed.Seed{
				seed.Seed{
						Name: "CreateJane",
						Run: func(db *gorm.DB) error {
						    return CreatePayment(db, "Boleto Bancário")
						},
				},
				seed.Seed{
						Name: "CreateJohn",
						Run: func(db *gorm.DB) error {
						    return CreatePayment(db, "Dinheiro")
						},
				},
		}
}
