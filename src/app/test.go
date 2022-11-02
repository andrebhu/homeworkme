package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)

func main() {
	// data := []byte(`{
	// 	"person": {
	// 	"name": {
	// 		"first": "Leonid",
	// 		"last": "Bugaev",
	// 		"fullName": "Leonid Bugaev"
	// 	},
	// 	"github": {
	// 		"handle": "buger",
	// 		"followers": 109
	// 	},
	// 	"avatars": [
	// 		{ "url": "https://avatars1.githubusercontent.com/u/14009?v=3&s=460", "type": "thumbnail" }
	// 	]
	// 	},
	// 	"company": {
	// 	"name": "Acme"
	// 	}
	// }`)

	data := []byte(`{
		"person": "Bob",
		"person": "Joe
	}`)

	value, err := jsonparser.GetString(data, "person")
	if err != nil {
		panic(err)
	}
	fmt.Printf(value)
}
