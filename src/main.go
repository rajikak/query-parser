package main

import (
	"fmt"
	"query-parser/parser"
)

func main() {

	// using the api
	p := parser.New("?filter=equals(lastName,'Smith')")

	err := p.Parse()
	if err != nil {
		fmt.Println(err)
	}
}
