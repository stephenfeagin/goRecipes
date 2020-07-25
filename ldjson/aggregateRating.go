package ldjson

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	kb "github.com/stephenfeagin/kitchenbox"
)

// unmarshalAggregateRating returns a *kb.AggregateRating from a *json.RawMessage, allowing for
// more flexible unmarshalling depending on the types used in the input JSON
func unmarshalAggregateRating(raw json.RawMessage) (*kb.AggregateRating, error) {
	// initialize a json.SyntaxError var for error checking
	var syntaxError *json.SyntaxError
	var err error

	agr := &kb.AggregateRating{}
	// Try to unmarshal into an aggregateRating struct. If successful, return
	if err = json.Unmarshal(raw, agr); err == nil {
		return agr, nil
	} else if errors.Is(err, syntaxError) {
		// If there's a JSON syntax error, there's nothing we can do
		return nil, err
	}

	// If unsuccessful, unmarshal into an empty raw aggregate rating struct
	rawAgr := &struct {
		Type         string          `json:"@type"`
		RatingValue  json.RawMessage `json:"ratingValue"`
		RatingCount  json.RawMessage `json:"ratingCount"`
		ItemReviewed string          `json:"itemReviewed"`
		BestRating   json.RawMessage `json:"bestRating"`
		WorstRating  json.RawMessage `json:"worstRating"`
	}{}

	// If unable to unmarshal into rawAgr, return SyntaxError if applicable, or else return the
	// empty struct (indicating empty input)
	if err = json.Unmarshal(raw, rawAgr); errors.Is(err, syntaxError) {
		return nil, err
	} else if err != nil {
		log.Printf("rawAgr unmarshal: %v\n", err)
		return agr, err
	}

	// Then parse the various strings into the result aggregateRating struct
	// First, copy over the existing string fields
	agr.Type, agr.ItemReviewed = rawAgr.Type, rawAgr.ItemReviewed

	// For numeric values, unmarshal into interface. Try to convert to their go numeric types, and
	// if unsuccessful, convert to string and then to the numeric type
	// RatingValue -- float64
	var valFloat float64
	err = json.Unmarshal(rawAgr.RatingValue, &valFloat)
	if err != nil {
		if errors.Is(err, syntaxError) {
			return nil, err
		}
		var valString string
		err = json.Unmarshal(rawAgr.RatingValue, &valString)
		if err != nil {
			return nil, err
		}
		valFloat, err = strconv.ParseFloat(valString, 64)
		if err != nil {
			return nil, err
		}
	}
	agr.RatingValue = valFloat

	// RatingCount -- int
	var countInt int
	err = json.Unmarshal(rawAgr.RatingCount, &countInt)
	if err != nil {
		if errors.Is(err, syntaxError) {
			return nil, err
		}
		var countString string
		err = json.Unmarshal(rawAgr.RatingCount, &countString)
		if err != nil {
			return nil, err
		}
		countInt, err = strconv.Atoi(countString)
		if err != nil {
			return nil, err
		}
	}
	agr.RatingCount = countInt

	// BestRating -- int
	var brInt int
	err = json.Unmarshal(rawAgr.BestRating, &brInt)
	if err != nil {
		if errors.Is(err, syntaxError) {
			return nil, err
		}
		var brString string
		err = json.Unmarshal(rawAgr.BestRating, &brString)
		if err != nil {
			return nil, err
		}
		brInt, err = strconv.Atoi(brString)
		if err != nil {
			return nil, err
		}
	}
	agr.BestRating = brInt

	// WorstRating -- int
	var wrInt int
	err = json.Unmarshal(rawAgr.WorstRating, &wrInt)
	if err != nil {
		if errors.Is(err, syntaxError) {
			return nil, err
		}
		var wrString string
		err = json.Unmarshal(rawAgr.WorstRating, &wrString)
		if err != nil {
			return nil, err
		}
		wrInt, err = strconv.Atoi(wrString)
		if err != nil {
			return nil, err
		}
	}
	agr.WorstRating = wrInt

	return agr, nil
}
