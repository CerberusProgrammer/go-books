package main

import (
	"database/sql"
	"fmt"
	"go-books/src/controllers"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola Mundo")
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	fmt.Println("Connecting to database...")
	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	defer db.Close()

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
