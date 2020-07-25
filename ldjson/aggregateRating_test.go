package ldjson

import (
	"encoding/json"
	"testing"
)

func TestAggregateRating(t *testing.T) {
	var rawRatingBareNumbers json.RawMessage = []byte(`
{
	"@type": "AggregateRating",
	"ratingValue": 4.25,
	"ratingCount": 100,
	"itemReviewed": "RecipeName",
	"bestRating": 5,
	"worstRating": 1
}
`)

	var rawRatingStringNumbers json.RawMessage = []byte(`
{
	"@type": "AggregateRating",
	"ratingValue": "4.25",
	"ratingCount": "100",
	"itemReviewed": "RecipeName",
	"bestRating": "5",
	"worstRating": "1"
}
`)

	var want = AggregateRating{
		Type:         "AggregateRating",
		RatingValue:  4.25,
		RatingCount:  100,
		ItemReviewed: "RecipeName",
		BestRating:   5,
		WorstRating:  1,
	}
	tests := map[string]json.RawMessage{
		"bareNumbers":   rawRatingBareNumbers,
		"stringNumbers": rawRatingStringNumbers,
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := processAggregateRatingFromJSON(&input)
			if err != nil {
				t.Fatal(err)
			}
			if *got != want {
				t.Fatal("incorrectly parsed aggregateRating")
			}
		})
	}
}
