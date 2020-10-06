package lexer

import (
	"fmt"
	"os"
)

// Lexer ...
type Lexer struct {
	data string

	buffer string
	index  int

	isEOF bool // To stop parsing

	Token           string
	TokenIdentifier int
}

// NewLexer ...
func NewLexer(code string) *Lexer {
	lexer := Lexer{
		data:   code,
		buffer: string(code[0]),
		index:  0,
		isEOF:  false,
	}
	lexer.NextToken()

	return &lexer
}

func (l *Lexer) err(message string) {
	fmt.Println("[LEXER] Error: " + message)
	os.Exit(1)
}

func (l *Lexer) next() {
	if l.index+1 >= len(l.data) {
		l.isEOF = true
		l.buffer = ""
		return
	}
	l.index++
	l.buffer = string(l.data[l.index])
}

// NextToken ...
func (l *Lexer) NextToken() {
	if l.isEOF {
		l.Token = ""
		l.TokenIdentifier = TokenEOF
		return
	}

	// Skipping spaces
	for isSpace(l.buffer) {
		l.next()
	}

	ident := l.buffer
	l.next()

	// Variable or word-operator
	if isWord(ident) || ident == "_" {
		for isWord(l.buffer) || isDigit(l.buffer) || l.buffer == "_" {
			ident += l.buffer
			l.next()
		}

		if val, ok := Tokens[ident]; ok {
			l.Token = ident
			l.TokenIdentifier = val
			return
		}

		l.Token = ident
		l.TokenIdentifier = TokenIdentifier
		return
	}

	// Is number
	if isDigit(ident) {
		for isDigit(l.buffer) {
			ident += l.buffer
			l.next()
		}

		// Float number
		if l.buffer == "." {
			ident += l.buffer
			l.next()

			for isDigit(l.buffer) {
				ident += l.buffer
				l.next()
			}
		}

		if ident[len(ident)-1] == '.' || isWord(l.buffer) {
			l.err("Unknown identifier after number definition (" + ident + l.buffer + ")")
		}

		l.Token = ident
		l.TokenIdentifier = TokenNumber
		return
	}

	// Symbols and other operators
	if val, ok := Tokens[ident]; ok {
		l.Token = ident
		l.TokenIdentifier = val
		return
	}

	l.err("Unknown identifier (" + ident + ")")
}
