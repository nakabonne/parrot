package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	var port = flag.String("p", "8080", "Set Port")
	flag.Parse()
	http.HandleFunc("/", handler)

	log.Println("The server is now ready to accept connections on port " + *port)

	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
