package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("p", "8080", "Set Port")
	flag.Parse()
	http.HandleFunc("/", request)

	log.Println("The server is now ready to accept connections on port " + *port)

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
