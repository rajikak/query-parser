package lexer

import (
	"testing"
)

func TestNextTokenOp(t *testing.T) {
	tests := []string{
		"?filter=equals(lastName,'Smith')",
		"?filter=lessThan(age,'25')",
		"?filter=lessOrEqual(lastModified,'2001-01-01')",
		"?filter=greaterThan(duration,'6:12:14')",
		"?filter=greaterOrEqual(percentage,'33.33')",
		"?filter=contains(description,'cooking')",
		"?filter=startsWith(description,'The')",
		"?filter=endsWith(description,'End')",
		"?filter=any(chapter,'Intro','Summary','Conclusion')",
		"?filter=has(articles)",
		"?filter=not(equals(lastName,null))",
		"?filter=or(has(orders),has(invoices))",
		"?filter=and(has(orders),has(invoices))",
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
		"filter=any(chapter,'Intro','Summary','Conclusion')",
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
