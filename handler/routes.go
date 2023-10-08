package handler

import (
	"UltravioletTest/dao"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func SetRoutes(router *mux.Router, db gorm.DB) {
	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		Books(&db, w)
	}).Methods("GET")

	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]
		Book(&db, w, bookID)
	}).Methods("GET")

	router.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		var book dao.Book
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		AddBook(&db, w, book)
	}).Methods("POST")

	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]
		var book dao.Book
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		UpdateBook(&db, w, book, bookID)
	}).Methods("PUT")

	router.HandleFunc("/books/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bookID := vars["id"]
		DeleteBook(&db, w, bookID)
	}).Methods("DELETE")

}
