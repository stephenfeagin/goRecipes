package ldjson

import (
	"encoding/json"
	"errors"

	kb "github.com/stephenfeagin/kitchenbox"
)

// unmarshalNutrition  attempts to unmarshal the raw data into a `nutrition` struct.
// If it's not successful, it returns the zero-value struct.
func unmarshalNutrition(raw json.RawMessage) (*kb.Nutrition, error) {
	var syntaxError *json.SyntaxError
	// Try to unmarshal into a kb.Nutrition struct. If not possible, assume that the input JSON
	// is empty, and return an empty kb.Nutrition struct.
	nut := &kb.Nutrition{}
	if err := json.Unmarshal(raw, nut); errors.Is(err, syntaxError) {
		return nil, err
	}

	return nut, nil
}
