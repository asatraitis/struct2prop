# Struct2Prop

[![License](https://img.shields.io/github/license/asatraitis/struct2prop)](LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/asatraitis/struct2prop)](https://goreportcard.com/report/github.com/asatraitis/struct2prop)

Utility that uses reflect std library to insepct and create JSON schema compatible properties from provided struct. Returned struct can be marshaled to a JSON schema for LLM function call consumption. Useful when building Go tooling around LLM's.

## Table of Contents

- [Installation](#installation)
- [Examples](#examples)

## Installation

```sh
go get github.com/asatraitis/struct2prop
```

## Examples

```go
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
```

Stdout:

```
{
  "type": "object",
  "properties": {
    "x": {
      "type": "number",
      "description": "horizontal value"
    },
    "y": {
      "type": "number",
      "description": "vertical value"
    },
    "z": {
      "type": "array",
      "description": "slice",
      "items": {
        "type": "string"
      }
    }
  }
}
```
