package seeds

import (
    "github.com/Zephiros/amarlinda/models"
    "gorm.io/gorm"
    "errors"
)

func CreatePayment(db *gorm.DB, name string) error {
    payment := models.Payment{Name: name}
    err := db.Select("payments.*").Where("payments.name = ?", name).First(&payment).Error
    if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
        return db.Create(&payment).Error
    }

    return err
}
