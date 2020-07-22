package json

import (
	"encoding/json"
)

// processNutritionFromJSON attempts to unmarshal the raw data into a `nutrition` struct.
// If it's not successful, it returns the zero-value struct.
func processNutritionFromJSON(raw *json.RawMessage) (*nutrition, error) {
	nut := &nutrition{}
	json.Unmarshal(*raw, nut)
	return nut, nil
}
