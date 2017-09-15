package main

import (
	"fmt"
	"log"
	"net/http"
)

func request(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("parse error!!!!", err)
	}

	//word := strings.Join(r.Form["word"], "")
	fmt.Println(r.Form)

}
