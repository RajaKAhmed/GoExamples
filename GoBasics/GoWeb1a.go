// Server1 is a minimal "echo" server.
package main
import (
"fmt"
"log"
"net/http"
"os"
)
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	port:=os.Getenv("PORT")
	if port == "" {
	 	port = "8081"
		log.Printf("Defaulting to port %s", port)
}
log.Printf("Listening on port %s", port)
if err := http.ListenAndServe(":"+port, nil); err != nil {
	log.Fatal(err)
}
}
// handler echoes the Path component of the requested URL.

func handler(w http.ResponseWriter, r *http.Request) {
/* 	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
} */
	fmt.Println("Enter Any text after the URL in the Address Bar")
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}