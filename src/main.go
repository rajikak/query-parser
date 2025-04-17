package main

import (
	"fmt"
	"query-parser/parser"
)

func main() {

	// using the api
	p := parser.New("?filter=equals(lastName,'Smith')")

	result, err := p.Parse()
	if err != nil {
		fmt.Println(err)
	} else {
		for _, filter := range result.Filters() {
			fmt.Println(filter.String())
		}
	}
}
