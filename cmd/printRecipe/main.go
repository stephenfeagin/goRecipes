package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/stephenfeagin/kitchenbox/ldjson"
)

func main() {
	var url, fname string
	urlCommand := flag.NewFlagSet("url", flag.ExitOnError)
	fileCommand := flag.NewFlagSet("file", flag.ExitOnError)
	urlCommand.StringVar(&url, "url", "", "URL to retrieve recipe from")
	fileCommand.StringVar(&fname, "file", "", "File to retrieve recipe from")

	if len(os.Args) < 2 {
		fmt.Println("url or file subcommand is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "url":
		urlCommand.Parse(os.Args[2:])
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			fmt.Println("request returned non-200 status")
			os.Exit(1)
		}

		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rec, err := ldjson.RetrieveRecipe(doc)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyJSON(rec)
	case "file":
		fileCommand.Parse(os.Args[2:])
		r, err := os.Open(fname)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer r.Close()
		doc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		rec, err := ldjson.RetrieveRecipe(doc)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		prettyJSON(rec)
	default:
		fmt.Println("url or file subcommand is required")
		os.Exit(1)
	}

}

func prettyJSON(i interface{}) error {
	pretty, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", string(pretty))
	return nil
}
