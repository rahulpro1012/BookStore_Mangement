package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/mux"

	// "github.com/gorilla/sessions"
	"net/http"
	"strconv"

	"github.com/rahul/go-bookstore/pkg/models"
	"github.com/rahul/go-bookstore/pkg/utils"
	"github.com/redis/go-redis/v9"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request, rdb *redis.Client, ctx context.Context) {
	// Check if books data exists in Redis cache
	cachedBooks, err := rdb.Get(ctx, "all_books").Result()
	if err == nil {
		// If data exists in cache, return cached data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(cachedBooks))
		return
	}

	// If data doesn't exist in cache, fetch it from the database
	newBooks := models.GetAllBooks()
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store the fetched data in Redis cache
	err = rdb.Set(ctx, "all_books", string(res), time.Second*30).Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
	}

	if ID < 0 {
		http.Error(w, "Book ID cannot be negative", http.StatusBadRequest)
		return
	}

	bookDetails, _, err := models.GetBookById(ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if ID < 0 {
		http.Error(w, "Book ID cannot be negative", http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Println("error while parsing")
	}
	book, err := models.DeleteBook(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if ID < 0 {
		http.Error(w, "Book ID cannot be negative", http.StatusBadRequest)
		return
	}
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db, err := models.GetBookById(ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
