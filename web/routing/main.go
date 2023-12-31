package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
** Get a book title from route variable and display it
 */
func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Getting the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

/*
** Create a new book
 */
func AddBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Adding the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

/*
** Update a book
 */
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Updating the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

/*
** Delete a book
 */
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Deleting the book...\n")
	fmt.Fprintf(w, "Title: %s\n", vars["title"])
}

/*
** List all books
 */
func ListBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listing books...\n")
}

func main() {
	r := mux.NewRouter()

	// A route with two variables
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	// Books path prefix and subrouter
	bookRouter := r.PathPrefix("/books").Subrouter()

	// Following route match only books/
	bookRouter.HandleFunc("/", ListBooks).Methods("GET")

	// Following route is added to handle /books
	bookRouter.HandleFunc("", ListBooks).Methods("GET")

	bookRouter.HandleFunc("/{title}", GetBook).Methods("GET")
	bookRouter.HandleFunc("/{title}", AddBook).Methods("POST")
	bookRouter.HandleFunc("/{title}", UpdateBook).Methods("PUT")
	bookRouter.HandleFunc("/{title}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(":8080", r)
}
