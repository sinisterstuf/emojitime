package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
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

// Emojify converts a time value into the "clock emoji" it represents.  Because
// there is only a clock emoji for every half an hour, the time is rounded to
// the nearest half hour.
func Emojify(t time.Time) string {
	var hs, ms string
	h := t.Hour()
	m := t.Minute()

	// round minutes
	switch {
	case 15 <= m && m < 45:
		ms = "30" // round up to half hour
	case 45 <= m:
		h++ // bump up hour
	}

	// normalise hours (there are no 24-hour emoji clocks)
	if h > 12 {
		h = h - 12
	}
	if h == 0 {
		h = 12
	}
	hs = strconv.Itoa(h)

	return ":clock" + hs + ms + ":"
}
