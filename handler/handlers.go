package handler

import (
	"UltravioletTest/dao"
	database "UltravioletTest/db"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Books(db *gorm.DB, w http.ResponseWriter) {
	fmt.Println("GET ALL BOOKS")
	bookList, err := database.GetBookList(db)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(bookList)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonData))
}

func Book(db *gorm.DB, w http.ResponseWriter, bookID string) {
	bookList, err := database.GetBook(db, bookID)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(bookList)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonData))
}

func AddBook(db *gorm.DB, w http.ResponseWriter, book dao.Book) {
	bookList, err := database.AddBook(db, book)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(bookList)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonData))
}

func UpdateBook(db *gorm.DB, w http.ResponseWriter, book dao.Book, bookID string) {
	bookList, err := database.UpdateBook(db, book, bookID)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	jsonData, err := json.Marshal(bookList)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(jsonData))
}

func DeleteBook(db *gorm.DB, w http.ResponseWriter, bookID string) {
	err := database.DeleteBook(db, bookID)
	if err != nil {
		fmt.Println("Error:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "Deleted")
}
