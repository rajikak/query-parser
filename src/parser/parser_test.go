package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := []string{
		"?filter=equals(lastName,'Smith')",
	}

	for _, test := range tests {
		p := New(test)
		err := p.Parse()

		if err != nil {
			t.Fatalf("test failure: %s", err)
		}
	}
}
