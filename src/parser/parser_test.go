package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []string{
		"?filter=equals(lastName,'Smith')",
		"?filter=lessThan(age,'25')",
		"?filter=lessOrEqual(lastModified,'2001-01-01')",
		"?filter=any(chapter,'Intro','Summary','Conclusion')",
		"?filter=has(articles)",
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
