package parser

import (
	"fmt"
	"query-parser/lexer"
	"slices"
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

func (p *Parser) Parse() error {
	err := p.match(lexer.ParamStartOp)
	if err != nil {
		return fmt.Errorf("syntax error: should start with %s", lexer.ParamStartOp)
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
	}
	return nil
}

func (p *Parser) nextToken() lexer.Token {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
	return p.curToken
}

func (p *Parser) checkToken(typ lexer.TokenType) bool {
	return p.curToken.Type == typ
}

func (p *Parser) checkPeek(typ lexer.TokenType) bool {
	return p.peekToken.Type == typ
}

func (p *Parser) match(typ lexer.TokenType) error {
	if !p.checkToken(typ) {
		return fmt.Errorf("syntax error: expected: %s, got: %s", typ, p.curToken)
	}
	return nil
}

func (p *Parser) function(tok string) bool {
	return slices.Contains([]string{lexer.LessThan, lexer.LessOrEqual, lexer.GreaterThan, lexer.GreaterOrEqual, lexer.Contains, lexer.Contains, lexer.StartWith, lexer.EndsWith, lexer.And, lexer.Has, lexer.Not, lexer.Or, lexer.And}, tok)
}

func (p *Parser) handleFilter() error {
	for {
		tok := p.nextToken()

		if tok.Type == lexer.Illegal {
			// syntax error
			return fmt.Errorf("syntax error: illegal chracter received:%s", tok)
		}

		if tok.Type == lexer.EndOfInput {
			if !p.reachedEnd() {
				return fmt.Errorf("syntax error: not ")
			}
			break
		}

		if !p.function(tok.Literal) {
			return fmt.Errorf("syntax error: function defintion is required, got: %s", tok)
		}

	}

	return nil
}

func (p *Parser) handleInclude() {
	fmt.Println("handle include")
}
