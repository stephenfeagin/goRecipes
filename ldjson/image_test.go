package ldjson

import (
	"encoding/json"
	"testing"

	kb "github.com/stephenfeagin/kitchenbox"
)

func TestParseImageJSON(t *testing.T) {
	var rawString json.RawMessage = []byte(`"https://blah.com/img.jpeg"`)
	var rawJSON json.RawMessage = []byte(`
		{
			"@type": "ImageObject",
			"url": "https://blah.com/img.jpeg",
                        "width": null,
                        "height": null,
                        "caption": null
		}
	`)
	want := kb.Image{
		Type: "ImageObject",
		URL:  "https://blah.com/img.jpeg",
	}
	tests := map[string]json.RawMessage{
		"rawURLString": rawString,
		"rawJSON":      rawJSON,
	}
	for name, input := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := unmarshalImage(input)
			if err != nil {
				t.Fatal(err)
			}
			if *got != want {
				t.Fatal("incorrectly processed Image")
			}
		})
	}
}
