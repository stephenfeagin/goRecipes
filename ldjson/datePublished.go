package ldjson

import (
	"regexp"
	"time"
)

// unmarshalDatePublished converts a raw string into a time.Time, allowing for inputs that do not
// include time/timezone data. If the input cannot be successfully parsed, it returns 1970-01-01
// and an error.
func unmarshalDatePublished(raw string) (time.Time, error) {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if re.MatchString(raw) {
		raw += "T00:00:00.00Z"
	}
	datePublished, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), err
	}
	return datePublished, nil
}
