package json

import (
	"encoding/json"
	"fmt"
)

func prettyPrint(p interface{}) {
	p, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return
	}
	fmt.Println(p)
}
