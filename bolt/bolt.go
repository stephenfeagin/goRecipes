// Package bolt provides CRUD functionality for a BoltDB datastore
package bolt

import (
	"encoding/json"
	"time"

	kb "github.com/stephenfeagin/kitchenbox"

	bolt "go.etcd.io/bbolt"
)

// CreateDB creates a boltDB file at the provided path if one does not already exist, and opens a
// connection to it.
func CreateDB(path string) (*bolt.DB, error) {
	db, err := bolt.Open(path, 0666, &bolt.Options{Timeout: 1 * time.Second})
	return db, err
}

// CloseDB closes the provided boltDB connection
func CloseDB(db *bolt.DB) {
	db.Close()
}

// CreateBuckets creates buckets with the provided names in the provided boltDB database
func CreateBuckets(db *bolt.DB, bucketNames []string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		var innerError error
		for _, name := range bucketNames {
			_, innerError = tx.CreateBucketIfNotExists([]byte(name))
		}
		return innerError
	})
	return err
}

func InsertRecipe(db *bolt.DB, bucket string, recipe *kb.Recipe) error {
	val, err := json.Marshal(recipe)
	if err != nil {
		return err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(recipe.Name), []byte(val))
		if err != nil {

			return err
		}
		return nil
	})

	return err
}

func GetRecipe(db *bolt.DB, bucket, name string) (*kb.Recipe, error) {
	var recipeBytes []byte
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		val := b.Get([]byte(name))
		recipeBytes = make([]byte, len(val))
		copy(recipeBytes, val)
		return nil
	})

	if err != nil {
		return nil, err
	}

	recipe := &kb.Recipe{}
	json.Unmarshal(recipeBytes, recipe)
	if err != nil {
		return nil, err
	}

	return recipe, nil
}

func ListAllKeys(db *bolt.DB, bucket string) ([]string, error) {
	var keys []string
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		c := b.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			keys = append(keys, string(k))
		}
		return nil
	})
	return keys, nil
}
