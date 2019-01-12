package main

import (
    "github.com/Autoblocks/go-dsl"
)

func NewTokenSet() dsl.TokenSet{
    return dsl.NewTokenSet(
		"LITERAL",
		"PLUS",
		"MINUS",
		"MULTIPLY",
		"DIVIDE",
		"OPEN_PAREN",
		"CLOSE_PAREN",
		"ASSIGN",
		"VARIABLE",
		"COMMENT",
		"NL",
		"EOF",
	)
}
