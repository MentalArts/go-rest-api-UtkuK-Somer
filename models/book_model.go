package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string   `json:"title"`
	AuthorID        uint     `json:"author_id"`
	Author          Author   `json:"author" gorm:"foreignKey:AuthorID"`
	ISBN            string   `json:"isbn" gorm:"unique"`
	PublicationYear int      `json:"publication_year"`
	Description     string   `json:"description"`
	Reviews         []Review `json:"reviews" gorm:"foreignKey:BookID"`
}
