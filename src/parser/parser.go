package parser

import (
	"query-parser/lexer"
)

type Parser struct {
	l *lexer.Lexer

	currToken lexer.Token
	peekToken lexer.Token
}

func New(l *lexer.Lexer) Parser {
	return Parser{l, l.NextToken(), l.NextToken()}
}

func (p Parser) CheckToken(typ lexer.TokenType) bool {
	return p.currToken.Type == typ
}

func (p Parser) CheckPeek(typ lexer.TokenType) bool {
	return p.peekToken.Type == typ
}

func (p Parser) NextToken() lexer.Token {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
	return p.l.NextToken()
}
