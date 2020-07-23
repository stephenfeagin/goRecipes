package json

import (
	"encoding/json"
	"errors"
)

// Nutrition defines the data for a schema.org NutritionInformation type
type Nutrition struct {
	Type                  string `json:"@type"`               // NutritionInformation
	Calories              string `json:"calories"`            // "233 calories"
	CarbohydrateContent   string `json:"carbohydrateContent"` // "2.4 g"
	CholesterolContent    string `json:"cholesterolContent"`
	FatContent            string `json:"fatContent"`
	FiberContent          string `json:"fiberContent"`
	ProteinContent        string `json:"proteinContent"`
	SaturatedFatContent   string `json:"saturatedFatContent"`
	ServingSize           string `json:"servingSize"`
	SodiumContent         string `json:"sodiumContent"`
	SugarContent          string `json:"sugarContent"`
	TransFatContent       string `json:"transFatContent"`
	UnsaturatedFatContent string `json:"unsaturatedFatContent"`
}

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
