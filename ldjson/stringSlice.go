package ldjson

import (
	"encoding/json"
	"strings"
)

// parseStringSliceFromJSON converts a json.RawMessage, which could be either a string or an array of strings,
// possibly empty, into a Go []string.
func parseStringSliceFromJSON(raw *json.RawMessage) ([]string, error) {
	// try to unmarshal into a string slice
	var catSlice []string
	if err := json.Unmarshal(*raw, &catSlice); err == nil {
		return catSlice, nil
	}
	// if not successful, cast to string and separate by commas
	catString := string(*raw)
	if catString == "" || catString == `""` {
		return catSlice, nil
	}
	catSlice = strings.Split(catString, ",")
	return catSlice, nil
}
