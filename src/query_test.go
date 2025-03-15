package main

import (
	"fmt"
	"testing"
)

func TestFunctionString(t *testing.T) {

	tests := []Function{
		{EndsWithFn, []string{"description", "End"}},
		{AnyFn, []string{"Intro", "Summary", "Conclusion"}},
		{OrFn, []Function{
			{HasFn, []string{"orders"}},
			{HasFn, []string{"invoices"}},
		}},
	}

	for _, test := range tests {
		fmt.Println(test.String())
	}
}
