package main

import (
	"fmt"
	"log"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("parse error!!!!", err)
	}
	fmt.Println(r.Form)
}
