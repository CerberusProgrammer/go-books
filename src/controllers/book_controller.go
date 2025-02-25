package controllers

import (
	"encoding/json"
	"go-books/src/models"
	"go-books/src/stores"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := stores.GetBooks()
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	book, err := stores.GetBookByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	stores.CreateBook(book)
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)
	err := stores.UpdateBook(id, updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := stores.DeleteBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
