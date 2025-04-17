package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []string{
		"?filter=equals(lastName,'Smith')",
	}

	for _, test := range tests {
		p := New(test)
		result, err := p.Parse()

		if err != nil {
			t.Fatalf("test failure: %s", err)
		}

		for _, filter := range result.Filters() {
			fmt.Println(filter.String())
		}
	}
}
