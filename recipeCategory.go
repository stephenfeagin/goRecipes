package main

import (
	"encoding/json"
	"strings"
)

func parseRecipeCategoryFromJSON(raw *json.RawMessage) ([]string, error) {
	// try to unmarshal into a string slice
	var catSlice []string
	if err := json.Unmarshal(*raw, &catSlice); err == nil {
		return catSlice, nil
	}
	// if not successful, cast to string and separate by commas
	catString := string(*raw)
	catSlice = strings.Split(catString, ",")
	return catSlice, nil
}
