package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)

func main() {
	data := []byte(`{
		"test": 1,
		"test": 2
		}`)

	r, _ := jsonparser.GetInt(data, "test")
	fmt.Println(r)
}
