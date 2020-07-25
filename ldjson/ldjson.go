package ldjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

var schemaRegexp = regexp.MustCompile(`"@type"\s*:\s*"Recipe"`)

// retrieveSelectionFromResponse traverses a goquery.Document
// document tree and searches for a <script type="application/ld+json"> element that matches a
// regexp for a schema.org Recipe. If it cannot parse a goquery.Document from the response, or if
// no nodes match the search, it returns nil and an error. Otherwise it returns a
// *goquery.Selection and nil.
func retrieveSelection(doc *goquery.Document) (*goquery.Selection, error) {
	var ldJSON *goquery.Selection

	doc.Find(`script[type="application/ld+json"]`).EachWithBreak(
		func(i int, s *goquery.Selection) bool {
			if schemaRegexp.MatchString(s.Text()) {
				ldJSON = s
				return false
			}
			return true
		})

	if ldJSON == nil {
		return nil, fmt.Errorf("no ld+json node found")
	}
	return ldJSON, nil
}

// isValidSchema makes a copy of the input buffer, checks if the content of the input matches the
// regex for a Schema.org Recipe JSON, then returns the result of the check and the copy of the
// input io.Reader
func isValidSchema(r io.Reader) (bool, io.Reader) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		return false, nil
	}

	match := schemaRegexp.Match(content)
	readerCopy := bytes.NewReader(content)
	return match, readerCopy
}

// extractFromSingleton produces a rawRecipe struct from a goquery.Selection that has only one
// LD+JSON object
func extractFromSingleton(s *goquery.Selection) (*rawRecipe, error) {
	recipe := &rawRecipe{}
	if err := json.Unmarshal([]byte(s.Text()), &recipe); err != nil {
		return nil, err
	}
	return recipe, nil
}

// extractFromSlice produces a rawRecipe struct from a goquery.Selection that has more than one
// LD+JSON object
func extractFromSlice(s *goquery.Selection) (*rawRecipe, error) {
	// unmarshal the text into an empty interface
	var jsonArray []interface{}

	if err := json.Unmarshal([]byte(s.Text()), &jsonArray); err != nil {
		return nil, err
	}

	recipeContainer := &rawRecipe{}
	for _, i := range jsonArray {
		// try to unmarshal into a Recipe struct
		obj, err := json.Marshal(i)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(obj, &recipeContainer); err != nil {
			continue
		} else {
			break
		}
	}
	if recipeContainer == nil {
		return nil, fmt.Errorf("couldn't parse JSON")
	}

	return recipeContainer, nil
}

// extractRecipe produces a rawRecipe struct from a goquery.Selection, trying both the
// extractFromSingle and the extractFromSlice approaches
func extractRecipe(s *goquery.Selection) (*rawRecipe, error) {
	var recipe *rawRecipe
	var err error
	recipe, err = extractFromSingleton(s)
	if err == nil {
		return recipe, nil
	}
	recipe, err = extractFromSlice(s)
	if err == nil {
		return recipe, nil
	}
	return nil, fmt.Errorf("extractRecipe: could not unmarshal JSON into a Recipe struct")
}
