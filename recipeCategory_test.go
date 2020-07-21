package main

import (
	"encoding/json"
	"testing"
)

func TestParseRecipeCategoryJSON(t *testing.T) {
	var (
		jsonArray      json.RawMessage = []byte(`["vegetarian","dinner"]`)
		stringSingle   json.RawMessage = []byte("vegetarian")
		stringMultiple json.RawMessage = []byte("vegetarian,dinner")
	)

	tests := []struct {
		name  string
		input json.RawMessage
		want  []string
	}{
		{"JSON Array", jsonArray, []string{"vegetarian", "dinner"}},
		{"String (single)", stringSingle, []string{"vegetarian"}},
		{"String (multiple)", stringMultiple, []string{"vegetarian", "dinner"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := parseRecipeCategoryFromJSON(&test.input)
			if err != nil {
				t.Fatal(err)
			}
			if len(got) != len(test.want) {
				t.Fatal("incorrectly parsed RecipeCategory: incorrect length")
			}
			for i := range got {
				if got[i] != test.want[i] {
					t.Fatal("incorrectly parsed RecipeCategory: incorrect contents")
				}
			}
		})
	}
}
