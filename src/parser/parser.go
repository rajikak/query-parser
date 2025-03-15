package parser

import (
	"go/token"
	"query-parser/lexer"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token

	//prefixParseFn map[token.Token]
}
