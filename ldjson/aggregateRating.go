package ldjson

import (
	"encoding/json"
	"errors"
	"strconv"

	kb "github.com/stephenfeagin/kitchenbox"
)

// unmarshalAggregateRating returns a *kb.AggregateRating from a *json.RawMessage, allowing for
// more flexible unmarshalling depending on the types used in the input JSON
func unmarshalAggregateRating(raw *json.RawMessage) (*kb.AggregateRating, error) {
	// initialize a json.SyntaxError var for error checking
	var syntaxError *json.SyntaxError

	agr := &kb.AggregateRating{}
	// Try to unmarshal into an aggregateRating struct. If successful, return
	if err := json.Unmarshal(*raw, agr); err == nil {
		return agr, nil
	} else if errors.Is(err, syntaxError) {
		// If there's a JSON syntax error, there's nothing we can do
		return nil, err
	}

	// If unsuccessful, unmarshal into an empty raw aggregate rating struct
	rawAgr := &struct {
		Type         string           `json:"@type"`
		RatingValue  *json.RawMessage `json:"ratingValue"`
		RatingCount  *json.RawMessage `json:"ratingCount"`
		ItemReviewed string           `json:"itemReviewed"`
		BestRating   *json.RawMessage `json:"bestRating"`
		WorstRating  *json.RawMessage `json:"worstRating"`
	}{}

	// If unable to unmarshal into rawAgr, return SyntaxError if applicable, or else return the
	// empty struct (indicating empty input)
	if err := json.Unmarshal(*raw, rawAgr); errors.Is(err, syntaxError) {
		return nil, err
	} else if err != nil {
		return agr, nil
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
