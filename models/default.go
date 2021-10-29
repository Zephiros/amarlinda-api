package models

import "time"

type Default struct {
	Id        uint      `json:"id" gorm:"primary_key;auto_increment"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;"`
}
