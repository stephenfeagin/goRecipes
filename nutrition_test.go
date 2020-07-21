package main

import (
	"encoding/json"
	"testing"
)

func TestParseNutritionJSON(t *testing.T) {
	var (
		emptyString     json.RawMessage = []byte("")
		emptyJSONObject json.RawMessage = []byte("{}")
		emptyJSONArray  json.RawMessage = []byte("[]")
		completeJSON    json.RawMessage = []byte(`
	{
		"@type": "NutritionInformation",
		"calories": "233 calories",
		"carbohydrateContent": "2.4 g",
		"cholesterolContent": "1 mg",
		"fatContent": "3 g",
		"fiberContent": "0 g",
		"proteinContent": "2 g",
		"saturatedFatContent": "1 g",
		"servingSize": "1 piece",
		"sodiumContent": "100 mg",
		"sugarContent": "8 g",
		"transFatContent": "0 g",
		"UnsaturatedFatContent": "2 g"
	}
    `)
	)

	emptyStruct := nutrition{}
	completeStruct := nutrition{
		Type:                  "NutritionInformation",
		Calories:              "233 calories",
		CarbohydrateContent:   "2.4 g",
		CholesterolContent:    "1 mg",
		FatContent:            "3 g",
		FiberContent:          "0 g",
		ProteinContent:        "2 g",
		SaturatedFatContent:   "1 g",
		ServingSize:           "1 piece",
		SodiumContent:         "100 mg",
		SugarContent:          "8 g",
		TransFatContent:       "0 g",
		UnsaturatedFatContent: "2 g",
	}
	tests := []struct {
		name  string
		input json.RawMessage
		want  nutrition
	}{
		{"Empty String", emptyString, emptyStruct},
		{"Empty JSON Object", emptyJSONObject, emptyStruct},
		{"Empty JSON Array", emptyJSONArray, emptyStruct},
		{"Complete JSON Object", completeJSON, completeStruct},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := processNutritionFromJSON(&test.input)
			if err != nil {
				t.Fatal(err)
			}
			if *got != test.want {
				t.Fatal("incorrectly processed Nutrition")
			}
		})
	}
}
