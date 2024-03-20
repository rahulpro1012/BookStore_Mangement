package routes

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/rahul/go-bookstore/pkg/controllers"
	"github.com/redis/go-redis/v9"
	"net/http"
)

var RegisterBookStoreRoutes = func(router *mux.Router, rdb *redis.Client, ctx context.Context) {

	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) { controllers.GetBook(w, r, rdb, ctx) }).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
