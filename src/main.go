package main

import "fmt"

func main() {

	// using the api
	p := New("?filter=equals(lastName,'Smith')")
	res := p.Parse()

	//for _, filter := range res.filters {
	//
	//}

	for _, include := range res.includes {
		fmt.Println(include.String())
	}
}
