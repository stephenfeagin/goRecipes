package main

import (
	"regexp"
	"time"
)

var re = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)

func parseDatePublishedFromJSON(raw string) (time.Time, error) {
	if re.MatchString(raw) {
		raw += "T00:00:00.00Z"
	}
	datePublished, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return time.Now(), err
	}
	return datePublished, nil
}
