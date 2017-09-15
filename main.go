package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", request)

	log.Println("Ready!")

	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
