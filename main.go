package main

import (
	"fmt"
	"log"
	"mentalartsapi/handlers"
	"mentalartsapi/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func waitForDB(dsn string, retries int, delay time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	for i := 0; i < retries; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ PostgreSQL bağlantısı başarılı!")
			return db, nil
		}

		fmt.Printf("❌ PostgreSQL bağlantısı başarısız! %d saniye sonra tekrar denenecek...\n", delay/time.Second)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("PostgreSQL bağlantısı başarısız: %v", err)
}

func main() {
	dsn := "host=db user=postgres password=123abcd dbname=mentalarts_db port=5432 sslmode=disable"

	db, err := waitForDB(dsn, 5, 5*time.Second)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	db.AutoMigrate(&models.Author{}, &models.Book{}, &models.Review{})

	handlers.InitDB(db)

	router := gin.Default()

	// 📌 Tüm API endpoint'lerini /api/v1 ile düzenleyelim
	api := router.Group("/api/v1")

	// 🟢 Author Routes
	api.POST("/authors", handlers.CreateAuthor)
	api.GET("/authors", handlers.GetAllAuthors)
	api.GET("/authors/:id", handlers.GetAuthor)
	api.PUT("/authors/:id", handlers.UpdateAuthor)
	api.DELETE("/authors/:id", handlers.DeleteAuthor)

	// 🟢 Book Routes
	api.POST("/books", handlers.CreateBook)
	api.GET("/books", handlers.GetAllBooks)
	api.GET("/books/:id", handlers.GetBook)
	api.PUT("/books/:id", handlers.UpdateBook)
	api.DELETE("/books/:id", handlers.DeleteBook)

	// 🟢 Review Routes
	api.POST("/books/:id/reviews", handlers.CreateReview)
	api.GET("/books/:id/reviews", handlers.GetReviewsByBook)
	api.PUT("/reviews/:id", handlers.UpdateReview)
	api.DELETE("/reviews/:id", handlers.DeleteReview)

	router.Run(":8000")
}
