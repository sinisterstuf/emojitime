package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestRootRoute(t *testing.T) {
	cases := []struct {
		method, body string
		status       int
	}{
		{"GET", "Emoji Timezone\n", 200},
		{"POST", "real logic here\n", 200},
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
