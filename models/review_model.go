package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	Rating     int       `json:"rating"`      // 1-5 arasında puan
	Comment    string    `json:"comment"`     // Kullanıcı yorumu
	DatePosted time.Time `json:"date_posted"` // İnceleme tarihi
	BookID     uint      `json:"book_id"`     // Foreign Key
	Book       Book      `json:"book" gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
}
