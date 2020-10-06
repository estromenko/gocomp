package parser

import (
	"fmt"
	"gocomp/compiler/lexer"
	"os"
)

// Parser ...
type Parser struct {
	lexer *lexer.Lexer
}

// NewParser ...
func NewParser(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer: lexer,
	}
}

func (p *Parser) error(message string) {
	fmt.Println("[PARSER] Error: " + message)
	os.Exit(1)
}
