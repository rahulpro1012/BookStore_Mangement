package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rahul/go-bookstore/pkg/routes"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("your-secret-key"))
var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer rdb.Close()

	r := mux.NewRouter()
	securedRoutes := r.PathPrefix("/api").Subrouter()
	securedRoutes.Use(middleWare)

	r.HandleFunc("/login", LoginHandler).Methods("GET")
	r.HandleFunc("/logout", LogoutHandler).Methods("GET")

	routes.RegisterBookStoreRoutes(securedRoutes, rdb, ctx)
	http.Handle("/", r)
	fmt.Println("server started at :9010")
	//log.Fatal(http.ListenAndServe("0.0.0.0:9010", r))
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

func middleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, "session-name")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if user is authenticated
		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Proceed to the next middleware or handler
		next.ServeHTTP(w, r)
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request /*, rdb *redis.Client*/) {
	// For demonstration purposes, let's assume the username is provided in the request query parameters.
	// In a real-world scenario, you would typically validate credentials from a login form or API request.

	session, err := store.Get(r, "session-name")
	if err != nil {
		fmt.Println("error while creating session", err)
	}
	session.Values["authenticated"] = true
	session.Options.MaxAge = 30
	session.Save(r, w)

	fmt.Println(session.Values)
	fmt.Fprintln(w, "Login Successful!")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		fmt.Println("error while ending session", err)
	}
	session.Values["authenticated"] = false
	session.Options.MaxAge = -1
	session.Save(r, w)
	fmt.Println(session.Values)
	fmt.Fprintln(w, "Logout Successful!")
}
