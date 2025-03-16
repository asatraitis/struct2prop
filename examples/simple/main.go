package main

import (
	"encoding/json"
	"fmt"

	"github.com/asatraitis/struct2prop"
)

type Location struct {
	X float32  `json:"x"description:"horizontal value"`
	Y float32  `json:"y"description:"vertical value"`
	Z []string `description:"slice"`
}

func main() {
	prop, err := struct2prop.GetProperties(Location{})
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(prop, "", "  ")
	fmt.Println(string(b))
}
