package main

import (
	"fmt"
	"strings"
)

type QueryParser struct {
	query string
}

type Function struct {
	name string
	args []any
}

func (f Function) String() string {
	msg := fmt.Sprintf("function: %s", f.name)

	var args []string
	for _, arg := range f.args {
		switch v := arg.(type) {
		case string:
			args = append(args, v)
		case Function:
			return v.String()
		}
	}

	return fmt.Sprintf("%s, args: [%s]", msg, strings.Join(args, ","))
}

const (
	EqualsFn         = "equals"
	LessThanFn       = "lessThan"
	LessOrEqualFn    = "lessOrEqual"
	GreaterThanFn    = "greaterThan"
	GreaterOrEqualFn = "greaterOrEqual"
	ContainsFn       = "contains"
	CountFn          = "count"
	StartWithFn      = "startsWith"
	EndsWithFn       = "endsWith"
	AnyFn            = "any"
	HasFn            = "has"
	NotFn            = "not"
	OrFn             = "or"
	AndFn            = "and"
)

type Filter struct {
	functions []Function
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

func (q QueryParser) Parse() QueryResult {
	res := QueryResult{}

	return res
}

func New(query string) QueryParser {
	return QueryParser{query}
}
