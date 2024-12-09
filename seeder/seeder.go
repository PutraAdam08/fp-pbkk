package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint64
	Title       string
	PublishYear string
	ISBN        string
	Units       []Unit `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category    string
}

type User struct {
	gorm.Model
	Email    string `gorm:"size:64,index"`
	Password string `gorm:"size:255"`
	IsAdmin  bool
}

type Category struct {
	ID   uint64
	Name string
}

type Unit struct {
	gorm.Model
	ID       uint64
	Code     string
	Borrowed bool
	BookID   uint64
	Borrow   []*Borrow `gorm:"many2many:user_languages;"`
}

type Borrow struct {
	gorm.Model
	ID            uint
	BorrowingDate *time.Time
	Deadline      *time.Time
	RetrievalDate *time.Time
	Status        int
	Units         []*Unit `gorm:"many2many:borrow_unit;"`
}

func main() {
	// Initialize the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("fail to connect db")
	}

	// AutoMigrate the schema
	db.AutoMigrate(&Book{}, &Unit{}, &Borrow{}, &User{})

	// Create some categories
	categories := []Category{
		{Name: "Fiction"},
		{Name: "Non-Fiction"},
		{Name: "Science"},
		{Name: "History"},
		{Name: "Biography"},
	}
	db.Create(&categories)
	pw := "12345678"
	p, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	users := []User{
		{Email: "admin1@gmail.com", Password: string(p), IsAdmin: true},
		{Email: "user1@gmail.com", Password: string(p), IsAdmin: false},
	}
	db.Create(&users)

	// Prepare a slice of books for bulk creation
	books := make([]Book, 0, 25)
	for i := 1; i <= 25; i++ {
		books = append(books, Book{
			Title:       fmt.Sprintf("Book Title %d", i),
			PublishYear: fmt.Sprintf("202%d", i%10),
			ISBN:        fmt.Sprintf("ISBN-%04d", i),
			Category:    categories[(i%4)+1].Name, // Cycle through category IDs
		})
	}

	// Perform the bulk creation
	result := db.Create(&books)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Println("Inserted book records:")
	for _, book := range books {
		fmt.Printf("ID: %d, Title: %s, Category: %s\n", book.ID, book.Title, book.Category)
	}
}
