package controllers

import (
	"encoding/json"
	"go-books/src/models"
	"go-books/src/stores"
	"net/http"
	"strconv"
)

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	authors := stores.GetAuthors()
	json.NewEncoder(w).Encode(authors)
}

func GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	author, err := stores.GetAuthorByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(author)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	stores.CreateAuthor(author)
	json.NewEncoder(w).Encode(author)
}

func UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var updatedAuthor models.Author
	json.NewDecoder(r.Body).Decode(&updatedAuthor)
	err := stores.UpdateAuthor(id, updatedAuthor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(updatedAuthor)
}

func DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	err := stores.DeleteAuthor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
