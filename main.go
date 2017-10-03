package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// HTTP Server
func main() {
	http.HandleFunc("/", rootRoute)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

// HTTP Routing for /
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

// App logic to handle Slack requests
func emojiTimezone(w http.ResponseWriter, r *http.Request) {
	if r.PostFormValue("command") != "/time" {
		msg := fmt.Sprint("srz can't handle that Slack command")
		http.Error(w, msg, 400)
		return
	}

	// extract time from text
	text := r.PostFormValue("text")
	t, err := time.Parse("15:04", text)
	if err != nil {
		// TODO: improve this error message
		log.Println("Failed to parse %q as time: %q", text, err.Error())
		http.Error(w, fmt.Sprintf("invalid time: %q", t), 400)
		return
	}

	clock := Emojify(t)
	fmt.Println(clock)

	fmt.Fprint(w, "done\n")
}

// Converts time into the relevant clock emoji rounded to nearest half hour
func Emojify(t time.Time) string {
	switch {
	case 0 <= t.Minute() && t.Minute() < 15:
		log.Println("rounding down to hour")
		return ":clock12:"
	case 15 <= t.Minute() && t.Minute() < 45:
		log.Println("rounding to half hour")
		return ":clock1230:"
	case 45 <= t.Minute():
		log.Println("rounding up to hour")
		return ":clock1:"
	default:
		log.Println("wtf")
		return ":question:"
	}

}
