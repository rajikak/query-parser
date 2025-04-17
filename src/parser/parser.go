package parser

import (
	"fmt"
	"query-parser/lexer"
	"slices"
	"strings"
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

func (p *Parser) Parse() (QueryResult, error) {
	result := QueryResult{}

	err := p.match(lexer.ParamStartOp)
	if err != nil {
		return result, fmt.Errorf("syntax error: should start with %s", lexer.ParamStartOp)
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
					return result, err
				}
				//p.nextToken()

				filter, err := p.handleFilter()
				if err != nil {
					return result, err
				}
				result.filters = append(result.filters, filter)

			} else if tok.Literal == lexer.Include {
				tok = p.nextToken()
				err = p.match(lexer.AssignOp)
				if err != nil {
					return result, err
				}
				//p.nextToken()
				err = p.handleInclude()
				if err != nil {
					return result, err
				}
			} else {
				return result, fmt.Errorf("syntax error: unknown operator: %s", tok)
			}
		}
	}
	return result, nil
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

func (p *Parser) peekNext() lexer.Token {
	return p.peekToken
}

func (p *Parser) match(typ lexer.TokenType) error {
	if !p.checkToken(typ) {
		return fmt.Errorf("syntax error: expected: %s, got: %s", typ, p.curToken)
	}
	return nil
}

func (p *Parser) function(tok string) bool {
	return slices.Contains([]string{
		lexer.LessThan,
		lexer.LessOrEqual,
		lexer.GreaterThan,
		lexer.GreaterOrEqual,
		lexer.Contains,
		lexer.Contains,
		lexer.StartWith,
		lexer.EndsWith,
		lexer.And,
		lexer.Has,
		lexer.Not,
		lexer.Or,
		lexer.And}, tok)
}

func (p *Parser) reachedEnd() bool {
	return p.l.ReachedEnd()
}

type stack struct {
	storage []lexer.Token
}

func (s stack) push(e lexer.Token) {
	s.storage = append(s.storage, e)
}

func (s stack) pop() lexer.Token {
	n := len(s.storage) - 1
	e := s.storage[n]
	s.storage = s.storage[:n]
	return e
}

func (p *Parser) handleFilter() (Filter, error) {
	var st stack

	filter := Filter{}

	for {
		tok := p.nextToken()

		if tok.Type == lexer.Illegal {
			// syntax error
			return filter, fmt.Errorf("syntax error: illegal chracter received: %s", tok)
		} else if tok.Type == lexer.EndOfInput {
			if !p.reachedEnd() {
				return filter, fmt.Errorf("syntax error: encountered end of input while buffer is not fully consumed")
			}
			break
		} else if tok.Type == lexer.LeftParenthesis {
			st.push(tok)
		} else if tok.Type == lexer.RightParenthesis {
			e := st.pop()
			if e.Type != lexer.LeftParenthesis {
				return filter, fmt.Errorf("syntax error: expected (")
			}
		} else {
			if !filter.defined() && !p.function(tok.Literal) {
				return filter, fmt.Errorf("syntax error: function defintion is required, got: %s", tok)
			}
			fun := Function{name: tok.Literal}

			if p.function(p.peekNext().Literal) {
				// nested functions - e.g. ?filter=and(has(orders),has(invoices))
				filter2, err := p.handleFilter()
				if err != nil {
					return filter2, err
				}
			} else {
				fmt.Println("val: ", tok.Literal)
				filter.functions = append(filter.functions, fun)
			}
		}
	}

	return filter, nil
}

func (p *Parser) handleInclude() error {
	return fmt.Errorf("not implemented yet, implementation is similar to Parser#handleFilter()")
}

type Function struct {
	name string
	args any
}

func (f Function) String() string {
	msg := fmt.Sprintf("function: %s", f.name)

	switch v := f.args.(type) {
	case []string:
		return fmt.Sprintf("%s, args: [%s]", msg, strings.Join(v, ","))
	case []Function:
		var msg1 string
		for _, fun := range v {
			msg1 += " " + fun.String()
		}
		return msg + msg1
	default:
		return "unknown"
	}
}

type Filter struct {
	functions []Function
}

func (f Filter) defined() bool {
	for _, filter := range f.functions {
		return filter.name != ""
	}
	return false
}

func (f Filter) String() string {
	var msg []string
	for _, fun := range f.functions {
		msg = append(msg, fun.String())
	}
	return strings.Join(msg, ",")
}

type Include struct {
	fields []string
}

func (i Include) String() string {
	return fmt.Sprintf("include fields: %s", strings.Join(i.fields, ","))
}

type QueryResult struct {
	filters  []Filter
	includes []Include
}

func (r QueryResult) Filters() []Filter {
	return r.filters
}
