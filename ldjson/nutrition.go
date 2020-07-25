package ldjson

import (
	"encoding/json"
	"errors"
)

// processNutritionFromJSON attempts to unmarshal the raw data into a `nutrition` struct.
// If it's not successful, it returns the zero-value struct.
func processNutritionFromJSON(raw *json.RawMessage) (*Nutrition, error) {
	var syntaxError *json.SyntaxError
	nut := &Nutrition{}
	if err := json.Unmarshal(*raw, nut); errors.Is(err, syntaxError) {
		return nil, err
	}

	return nut, nil
}
