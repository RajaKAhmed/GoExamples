package main

import (
	"fmt"
	"net/http"
	"os"
)

func hostHandler(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()

	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, "<h1>HOSTNAME: %s</h1><br>", name)
	fmt.Fprintf(w, "<h1>ENVIRONMENT VARS:</h1><br>")
	fmt.Fprintf(w, "<ul>")


}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Awesome Website in Go</h1><br>")
	fmt.Fprintf(w, "<a href='/host/'>Host info</a><br>")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/host/", hostHandler)
	http.ListenAndServe(":8080", nil)
}
