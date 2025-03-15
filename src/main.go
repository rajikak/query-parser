package main

import "fmt"

func main() {

	// using the api
	p := New("?filter=equals(lastName,'Smith')")
	res := p.Parse()

	for _, filter := range res.filters {
		if filter.fn == EqualsFn {
			fmt.Println("look up resources by equality")
			filter.Print()
		}
	}

	for _, include := range res.includes {
		include.Print()
	}
}
