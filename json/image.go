package json

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Image defines the data for a schema.org Image. It only includes fields that are frequently
// included in Recipe models
type Image struct {
	Type    string `json:"@type"` // should be ImageObject
	URL     string `json:"url"`
	Width   string `json:"width"`
	Height  string `json:"height"`
	Caption string `json:"caption"`
}

// processRawImageFromJSON turns a *json.RawMessage into an Image struct. The RawMessage will either be a
// plain URL string or an actual ImageObject JSON representation according to schema.org
func processRawImageFromJSON(raw *json.RawMessage) (*Image, error) {

	// first try to unmarshal into an Image struct
	img := &Image{}
	err := json.Unmarshal(*raw, img)
	if err == nil {
		return img, nil
	}

	var syntaxError *json.SyntaxError
	if errors.Is(err, syntaxError) {
		return nil, err
	}
	// if that didn't work, then unmarshal into an empty interface
	var imgInterface interface{}
	json.Unmarshal(*raw, &imgInterface)
	// type assert into string
	imgString, ok := imgInterface.(string)
	if !ok {
		return nil, fmt.Errorf("couldn't assert into string")
	}
	img.Type = "ImageObject"
	img.URL = imgString

	return img, nil

}
