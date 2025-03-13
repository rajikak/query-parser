package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {

	tests := []string{
		"courses?filter=equals(",
		"students?filter=equals(displayName,",
		"courses?filter=equals(displayName,)",
		"students?filter=equals(displayName,null)",
		"teachers?filter=equals(displayName,lastName)",
		"courses?filter=equals(displayName,'Brian Connor')",
		"users?filter=equals(displayName,'Brian Connor')",
		"customers?filter=greaterThan(count(orders),count(invoices))",
		"blogs?filter=lessThan(count(owner.articles),'10')",
		"blogs?include=owner.articles.revisions&filter=and(or(equals(title,'Technology'),has(owner.articles)),not(equals(owner.lastName,null)))&filter[owner.articles]=equals(caption,'Two')&filter[owner.articles.revisions]=greaterThan(publishTime,'2005-05-05')",
	}

	for _, test := range tests {
		l := New(test)

		for {
			tok := l.NextToken()
			tok.Print()
			if tok.Type == EndOfInput || tok.Type == Illegal {
				break
			}
		}
	}
}
