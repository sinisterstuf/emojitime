package emoji

import (
	"testing"
	"time"
)

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
