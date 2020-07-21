package main

import (
	"testing"
	"time"
)

func TestParseDatePublishedJSON(t *testing.T) {
	want := time.Date(2019, 4, 2, 0, 0, 0, 0, time.UTC)
	tests := map[string]string{
		"Full Datetime": "2019-04-02T00:00:00.000Z",
		"Date Only":     "2019-04-02",
	}

	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			datePublished, err := parseDatePublishedFromJSON(input)
			if err != nil {
				t.Fatal(err)
			}
			if datePublished != want {
				t.Fatal("incorrectly parsed DatePublished")
			}
		})
	}
}
