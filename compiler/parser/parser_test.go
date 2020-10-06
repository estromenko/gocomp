package parser

import (
	"gocomp/compiler/lexer"
	"testing"
)

func TestParser(t *testing.T) {
	data := "var asd123"
	lex := lexer.NewLexer(data)
	_ = NewParser(lex)
}
