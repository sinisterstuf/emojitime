package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRootRoute(t *testing.T) {
	cases := []struct {
		method, body string
		status       int
	}{
		{"GET", "Emoji Timezone\n", 200},
		{"POST", "srz can't handle that Slack command\n", 400},
		{"PUSH", "Can't handle \"PUSH\" on route \"/\"!\n", 405},
	}
	for _, c := range cases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(c.method, "/", nil)
		rootRoute(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		got := struct {
			body   string
			status int
		}{
			string(body),
			resp.StatusCode,
		}

		if got.body != c.body || got.status != c.status {
			t.Errorf(
				"Request %q to / == %q (%d), want %q (%d)",
				c.method, got.body, got.status, c.body, c.status,
			)
		}
	}
}

func TestEmojify(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"00:00", ":clock12:"},
		{"12:00", ":clock12:"},
		{"12:05", ":clock12:"},
		{"12:14", ":clock12:"},
		{"12:15", ":clock1230:"},
		{"12:30", ":clock1230:"},
		{"12:44", ":clock1230:"},
		{"12:55", ":clock1:"},
		{"12:45", ":clock1:"},
		{"12:59", ":clock1:"},
		{"13:00", ":clock1:"},
	}

	for _, c := range cases {
		ti, _ := time.Parse("15:04", c.in)
		got := Emojify(ti)
		if got != c.want {
			t.Errorf("Emojify(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
