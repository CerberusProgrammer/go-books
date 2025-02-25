package main

import (
	"fmt"
	"go-books/src/controllers"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)

	http.HandleFunc("/authors", controllers.GetAuthors)
	http.HandleFunc("/author", controllers.GetAuthorByID)
	http.HandleFunc("/author/create", controllers.CreateAuthor)
	http.HandleFunc("/author/update", controllers.UpdateAuthor)
	http.HandleFunc("/author/delete", controllers.DeleteAuthor)

	http.HandleFunc("/books", controllers.GetBooks)
	http.HandleFunc("/book", controllers.GetBookByID)
	http.HandleFunc("/book/create", controllers.CreateBook)
	http.HandleFunc("/book/update", controllers.UpdateBook)
	http.HandleFunc("/book/delete", controllers.DeleteBook)

	http.ListenAndServe(":8080", nil)
}
