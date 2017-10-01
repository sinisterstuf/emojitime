package main

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestRootRoute(t *testing.T) {
	cases := []struct {
		method, want string
	}{
		{"GET", "Emoji Timezone"},
		{"POST", "real logic here"},
		{"PUSH", "some error message"},
	}
	for _, c := range cases {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(c.method, "/", nil)
		rootRoute(w, req)

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		got := string(body)

		if got != c.want {
			t.Errorf("Request %q to / == %q, want %q", c.method, body, c.want)
		}
	}
}
