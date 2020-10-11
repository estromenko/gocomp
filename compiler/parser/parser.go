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

func (p *Parser) err(message string) {
	fmt.Println("[PARSER] Error: " + message)
	os.Exit(1)
}

// Parse ...
func (p *Parser) Parse() *Node {
	node := NewNode(KindProgram, nil, p.statement(), nil, nil)

	if p.lexer.TokenIdentifier != lexer.TokenEOF {
		p.err("Invalid statement syntax")
	}

	return node
}

func (p *Parser) statement() *Node {
	if p.lexer.TokenIdentifier == lexer.TokenIf {
		_ = NewNode(KindIf, nil, p.parenExpresion(), nil, nil)

	}

	p.lexer.NextToken()
	return nil
}

func (p *Parser) parenExpresion() *Node {
	if p.lexer.TokenIdentifier != lexer.TokenLeftPar {
		p.err(`"(" expexted`)
	}
	p.lexer.NextToken()
	n := p.expresion()

	if p.lexer.TokenIdentifier != lexer.TokenRightPar {
		p.err(`")" expected`)
	}

	p.lexer.NextToken()
	return n
}

func (p *Parser) expresion() *Node {
	return nil
}
