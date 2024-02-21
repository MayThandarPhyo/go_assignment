package main

import (
	"log"
	"net/http"
)

func main() {
	// fmt.Println("Hello world!")
	http.Handle("/", http.FileServer(http.Dir("templates")))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err.Error())
	}
	
}