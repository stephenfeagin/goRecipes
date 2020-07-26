package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/PuerkitoBio/goquery"

	"github.com/stephenfeagin/kitchenbox/bolt"
	"github.com/stephenfeagin/kitchenbox/ldjson"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("need one command line argument")
	}
	var err error
	db, err := bolt.CreateDB("./store.db")
	if err != nil {
		log.Fatal(err)
	}
	defer bolt.CloseDB(db)
	if err := bolt.CreateBuckets(db, []string{"recipes"}); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(err)
	}
	recipe, err := ldjson.RetrieveRecipe(doc)
	if err != nil {
		log.Fatal(err)
	}

	err = bolt.InsertRecipe(db, "recipes", recipe)
	if err != nil {
		log.Fatal(err)
	}

	rec, err := bolt.GetRecipe(db, "recipes", recipe.Name)
	if err != nil {
		log.Fatal(err)
	}

	if !reflect.DeepEqual(recipe, rec) {
		log.Fatal("Not Equal!")
	}

	recJSON, err := json.MarshalIndent(rec, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(recJSON))

}
