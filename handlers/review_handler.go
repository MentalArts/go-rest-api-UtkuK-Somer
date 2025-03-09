package handlers

import (
	"mentalartsapi/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Yeni Review ekleme
func CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review.DatePosted = time.Now() // Tarihi otomatik ayarla

	result := db.Create(&review)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, review)
}

// Bir Kitaba Ait Tüm İncelemeleri Listeleme
func GetReviewsByBook(c *gin.Context) {
	bookID := c.Param("id") // Kitabın ID'sini al

	var reviews []models.Review

	if err := db.Preload("Book").Preload("Book.Author").Where("book_id = ?", bookID).Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// Belirli Bir İncelemeyi Güncelleme
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := db.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		return
	}

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, review)
}

// İncelemeyi Silme
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review models.Review

	if err := db.First(&review, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := db.Delete(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "review deleted"})
}
