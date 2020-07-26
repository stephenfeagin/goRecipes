package ldjson

import (
	"encoding/json"
	"errors"
	"log"

	kb "github.com/stephenfeagin/kitchenbox"
)

// unmarshalImage turns a *json.RawMessage into an Image struct. The RawMessage will either be a
// plain URL string or an actual ImageObject JSON representation according to schema.org
func unmarshalImage(raw json.RawMessage) (*kb.Image, error) {

	// First try to unmarshal into an Image struct
	// If there's a JSON syntax error, there's nothing we can do about it
	img := &kb.Image{}
	var syntaxError *json.SyntaxError
	if err := json.Unmarshal(raw, img); err == nil {
		return img, nil
	} else if errors.Is(err, syntaxError) {
		log.Println(err)
		return nil, err
	}

	var imgString string
	if err := json.Unmarshal(raw, &imgString); err != nil {
		return nil, err
	}
	img.Type = "ImageObject"
	img.URL = imgString

	return img, nil

}
