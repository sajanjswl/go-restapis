package main

import (
	"strconv"

	"github.com/gorilla/mux"

	"encoding/json"
	"log"
	"net/http"

	//"reflect"
	"fmt"
)

type Book struct {
	ID     int    `jason:id`
	Title  string `jason:title`
	Author string `jason:author`
	Year   string `jason:year`
}

var books []Book

func main() {
	fmt.Println("starting RestApi server !..")
	router := mux.NewRouter()
	books = append(books, Book{ID: 1, Title: "golang", Author: "Mr. Golang", Year: "2019"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2018"},
		Book{ID: 3, Title: "Channel", Author: "Mr. Channel", Year: "2019"},
		Book{ID: 4, Title: "Concourrency", Author: "Mr. Concurrency", Year: "2019"},
	)

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {

	_ = json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	//log.Println(params)

	// reflect give us the type of data type
	//fmt.Println(reflect.TypeOf(params["id"]))

	i, _ := strconv.Atoi(params["id"])

	//fmt.Println(reflect.TypeOf(i))

	for _, book := range books {
		if book.ID == i {
			_ = json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	_ = json.NewEncoder(w).Encode(books)

	fmt.Println(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}

	}

	_ = json.NewEncoder(w).Encode(books)

}

func removeBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	//fmt.Println(reflect.TypeOf(i))

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

	_ = json.NewEncoder(w).Encode(books)
}
