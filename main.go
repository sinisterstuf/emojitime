package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", rootRoute)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func rootRoute(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprint(w, "Emoji Timezone")
	case "POST":
		fmt.Fprint(w, "real logic here")
	default:
		fmt.Fprintf(w, "Can't handle %q on route %q!", r.Method, r.URL)
	}
}
