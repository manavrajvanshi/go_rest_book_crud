package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

//Book Model - Struct

type Book struct{
	ID 	string	`json:"id"`
	Isbn 	string 	`json:"isbn"`
	Title 	string 	`json:"title"`
	Author 	*Author `json:"author"`
}

// Author Model

type Author struct{
	FirstName	string	`json:"firstName"`
	LastName	string	`json:"lastName"`
}

// Init dummy book slice

var books []Book

//Get All Books
func getBooks( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

//Get Single Book
func getBook( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	//Get params
	params := mux.Vars(r)
	// Loop through books and find with Id

	for _, item := range books{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	fmt.Fprintf( w, "Sorry, Book Id %s does not exist", params["id"])
}

//Create a new Book
func createBook( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa( rand.Intn(10000000) )
	books = append( books, book )
//	fmt.Println("New Book ", book )
	json.NewEncoder(w).Encode(book)
}

//Update a Book
func updateBook( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	for index, item := range books{
		if item.ID == book.ID {
			books[index] = book
			books[index].ID = book.ID
			json.NewEncoder(w).Encode(books[index])
			return
		}
	}
	json.NewEncoder(w).Encode( &Book{} )
}

//Delete a book
func deleteBook( w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"]{
			books = append( books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}


func main() {

	// Init Router
	router := mux.NewRouter()

	// Mock Data for Books and Authors
	books = append( books, Book { ID : "1", Isbn : "11", Title : "Book 1", Author : &Author { FirstName : "Alice", LastName : "Doe" } } )
	books = append( books, Book { ID : "2", Isbn : "22", Title : "Book 2", Author : &Author { FirstName : "Bob", LastName : "Doe" } } )
	books = append( books, Book { ID : "3", Isbn : "33", Title : "Book 3", Author : &Author { FirstName : "Charles", LastName : "Doe" } } )
	books = append( books, Book { ID : "4", Isbn : "44", Title : "Book 4", Author : &Author { FirstName : "Devin", LastName : "Doe" } } )

	//Route Handlers / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal( http.ListenAndServe( ":3005", router ) )

}
