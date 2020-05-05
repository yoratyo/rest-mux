package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8080"

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server Up and Running..")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/post", addPost).Methods("POST")
	router.HandleFunc("/post/{id}", getPost).Methods("GET")

	log.Println("Server listening on port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
