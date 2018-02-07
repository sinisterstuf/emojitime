// Package emoji provides functions for converting various types into their
// emoji representations.
package emoji

import (
	"strconv"
	"time"
)

// Clockify converts a time value into the "clock emoji" it represents.
// Because there is only a clock emoji for every half an hour, the time is
// rounded to the nearest half hour.
func Clockify(t time.Time) string {
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

// Flagify converts a 2-letter country code into the "flag emoji" it represents.
// It assumes the input to be a valid ISO 3166-1 alpha-2 country code and does
// simple string concatenation to produce the output.
func Flagify(s string) string {
	return ":flag-" + s + ":"
}
