package ldjson

import (
	"encoding/json"
	"testing"

	kb "github.com/stephenfeagin/kitchenbox"
)

func TestUnmarshalAggregateRating(t *testing.T) {
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

	var want = kb.AggregateRating{
		Type:         "AggregateRating",
		RatingValue:  4.25,
		RatingCount:  100,
		ItemReviewed: "RecipeName",
		BestRating:   5,
		WorstRating:  1,
	}

	tests := []struct {
		name  string
		input json.RawMessage
		want  kb.AggregateRating
	}{
		{"Bare Numbers", rawRatingBareNumbers, want},
		{"String Numbers", rawRatingStringNumbers, want},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := unmarshalAggregateRating(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if *got != test.want {
				t.Fatal("incorrectly parsed AggregateRating")
			}
		})
	}
}
