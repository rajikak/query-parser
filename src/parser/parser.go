package parser

import (
	"fmt"
	"query-parser/lexer"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  lexer.Token
	peekToken lexer.Token
}

func New(input string) Parser {
	l := lexer.New(input)
	return Parser{l, l.NextToken(), l.NextToken()}
}

func (p Parser) Parse() error {
	err := p.match(lexer.ParamStartOp)
	if err != nil {
		return err // syntax error
	}
	for {

		tok := p.nextToken()
		if tok.Type == lexer.EndOfInput || tok.Type == lexer.Illegal {
			break
		}

		if tok.Type == lexer.Keyword {
			if tok.Literal == lexer.Filter {
				tok = p.nextToken()
				err = p.match(lexer.AssignOp)
				if err != nil {
					return err
				}
				p.nextToken()
				p.handleFilter()

			} else if tok.Literal == lexer.Include {
				tok = p.nextToken()
				err = p.match(lexer.AssignOp)
				if err != nil {
					return err
				}
				p.nextToken()
				p.handleInclude()
			} else {
				return fmt.Errorf("syntax error: unknown operator: %s", tok)
			}
		}

		//if tok.Type == lexer.Keyword {
		//	err = p.match(lexer.AssignOp)
		//	if err != nil {
		//		return err // syntax error
		//	}
		//	tok = p.nextToken()
		//
		//} else if tok.Type == lexer.Include {
		//	err = p.match(lexer.AssignOp)
		//	if err != nil {
		//		return err // syntax error
		//	}
		//}
	}

	return nil
}

func (p Parser) nextToken() lexer.Token {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
	return p.curToken
}

func (p Parser) checkToken(typ lexer.TokenType) bool {
	return p.curToken.Type == typ
}

func (p Parser) checkPeek(typ lexer.TokenType) bool {
	return p.peekToken.Type == typ
}

func (p Parser) match(typ lexer.TokenType) error {
	if !p.checkToken(typ) {
		return fmt.Errorf("syntax error: expected: %s, got: %s", typ, p.curToken)
	}
	return nil
}

func (p Parser) handleFilter() {
	fmt.Println("handle filter")
}

func (p Parser) handleInclude() {
	fmt.Println("handle include")
}
