package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/rahul/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB, error) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	if db.RecordNotFound() {
		return nil, db, errors.New("book not found")
	}
	return &getBook, db, nil
}

func DeleteBook(Id int64) (Book, error) {
	var book Book
	result := db.Where("ID=?", Id).Delete(&book)

	// Check if the record was not found
	if result.RowsAffected == 0 {
		return book, errors.New("book not found")
	}

	return book, nil
}
