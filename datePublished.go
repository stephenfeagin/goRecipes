package main

import (
	"regexp"
	"time"
)

func parseDatePublishedFromJSON(raw string) (time.Time, error) {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if re.MatchString(raw) {
		raw += "T00:00:00.00Z"
	}
	datePublished, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC), err
	}
	return datePublished, nil
}
