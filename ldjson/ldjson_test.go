package ldjson

import (
	"log"
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var testInputFiles = map[string]string{
	"Food Network":    "testdata/foodnetwork-ginger-salmon.html",
	"AllRecipes":      "testdata/allrecipes-turkey-burgers.html",
	"BonAppetit":      "testdata/bonappetit-brown-butter-peach-cobbler.html",
	"Serious Eats":    "testdata/seriouseats-margarita.html",
	"Cookie and Kate": "testdata/cookieandkate-marinated-chickpeas.html",
	"Kitchn":          "testdata/kitchn-strawberry-shortcake.html",
}

func init() {
	log.SetFlags(log.Lshortfile)
}

func TestRetrieveSelection(t *testing.T) {
	for name, input := range testInputFiles {
		t.Run(name, func(t *testing.T) {
			f, err := os.Open(input)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			doc, err := goquery.NewDocumentFromReader(f)
			if err != nil {
				t.Fatal(err)
			}

			if _, err = retrieveSelection(doc); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestRetrieveRecipe(t *testing.T) {
	for name, input := range testInputFiles {
		t.Run(name, func(t *testing.T) {

			f, err := os.Open(input)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			doc, err := goquery.NewDocumentFromReader(f)
			if err != nil {
				t.Fatal(err)
			}

			if _, err = RetrieveRecipe(doc); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func BenchmarkExtractRecipe(b *testing.B) {
	for name, input := range testInputFiles {
		b.Run(name, func(b *testing.B) {
			f, err := os.Open(input)
			if err != nil {
				b.Fatal(err)
			}
			defer f.Close()

			doc, err := goquery.NewDocumentFromReader(f)
			if err != nil {
				b.Fatal(err)
			}

			if _, err := RetrieveRecipe(doc); err != nil {
				b.Fatal(err)
			}
		})
	}
}
