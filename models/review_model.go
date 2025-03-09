package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating     int       `json:"rating"`
	Comment    string    `json:"comment"`
	DatePosted time.Time `json:"date_posted"`
	BookID     uint      `json:"book_id"`
	Book       Book      `json:"book" gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
}
