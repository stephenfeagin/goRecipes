package ldjson

import (
	"encoding/json"
	"errors"
	"fmt"

	kb "github.com/stephenfeagin/kitchenbox"
)

// unmarshalImage turns a *json.RawMessage into an Image struct. The RawMessage will either be a
// plain URL string or an actual ImageObject JSON representation according to schema.org
func unmarshalImage(raw *json.RawMessage) (*kb.Image, error) {

	// First try to unmarshal into an Image struct
	// If there's a JSON syntax error, there's nothing we can do about it
	img := &kb.Image{}
	var syntaxError *json.SyntaxError
	if err := json.Unmarshal(*raw, img); err == nil {
		return img, nil
	} else if errors.Is(err, syntaxError) {
		return nil, err
	}

	// If that didn't work, then unmarshal into an empty interface
	var imgInterface interface{}
	json.Unmarshal(*raw, &imgInterface)
	// type assert into string
	imgString, ok := imgInterface.(string)
	if !ok {
		return nil, fmt.Errorf("image: couldn't assert into string")
	}
	img.Type = "ImageObject"
	img.URL = imgString

	return img, nil

}
