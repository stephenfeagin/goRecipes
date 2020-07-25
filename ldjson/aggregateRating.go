package ldjson

import (
	"encoding/json"
	"errors"
	"strconv"
)

// rawAggregateRating is a more flexible representation of a schema.org AggregateRating object. It
// allows for number fields to be in string form.
type rawAggregateRating struct {
	Type         string           `json:"@type"` // AggregateRating
	RatingValue  *json.RawMessage `json:"ratingValue"`
	RatingCount  *json.RawMessage `json:"ratingCount"`
	ItemReviewed string           `json:"itemReviewed"` // should match allRecipes.Name
	BestRating   *json.RawMessage `json:"bestRating"`
	WorstRating  *json.RawMessage `json:"worstRating"`
}

func unmarshalAggregateRating(raw *json.RawMessage) (*AggregateRating, error) {
	// initialize a json.SyntaxError var for error checking
	var syntaxError *json.SyntaxError

	agr := &AggregateRating{}
	// Try to unmarshal into an aggregateRating struct
	if err := json.Unmarshal(*raw, agr); err == nil {
		return agr, nil
	} else if errors.Is(err, syntaxError) {
		// If there's a JSON syntax error, there's nothing we can do
		return nil, err
	}

	// If unsuccessful, unmarshal into rawAggregateRating
	rawAgr := &rawAggregateRating{}
	if err := json.Unmarshal(*raw, rawAgr); err != nil {
		return nil, err
	}
	// Then parse the various strings into the result aggregateRating struct
	// First, copy over the existing string fields
	agr.Type, agr.ItemReviewed = rawAgr.Type, rawAgr.ItemReviewed

	// Unmarshal the numeric fields (they're *json.RawMessage), then parse into
	// numeric types
	// RatingValue -- float64
	var valString string
	if err := json.Unmarshal(*rawAgr.RatingValue, &valString); err != nil {
		return nil, err
	}
	if valFloat, err := strconv.ParseFloat(valString, 64); err != nil {
		return nil, err
	} else {
		agr.RatingValue = valFloat
	}

	// RatingCount -- int
	var countString string
	if err := json.Unmarshal(*rawAgr.RatingCount, &countString); err != nil {
		return nil, err
	}
	if countInt, err := strconv.Atoi(countString); err != nil {
		return nil, err
	} else {
		agr.RatingCount = countInt
	}

	// BestRating -- int
	var brString string
	if err := json.Unmarshal(*rawAgr.BestRating, &brString); err != nil {
		return nil, err
	}
	if brInt, err := strconv.Atoi(brString); err != nil {
		return nil, err
	} else {
		agr.BestRating = brInt
	}

	// WorstRating -- int
	var wrString string
	if err := json.Unmarshal(*rawAgr.WorstRating, &wrString); err != nil {
		return nil, err
	}
	if wrInt, err := strconv.Atoi(wrString); err != nil {
		return nil, err
	} else {
		agr.WorstRating = wrInt
	}

	return agr, nil
}
