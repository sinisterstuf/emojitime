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
		fmt.Fprint(w, "Emoji Timezone\n")
	case "POST":
		emojiTimezone(w, r)
	default:
		msg := fmt.Sprintf("Can't handle %q on route %q!", r.Method, r.URL)
		http.Error(w, msg, 405)
	}
}

func emojiTimezone(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "real logic here\n")
}
