package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {

	data := "var _variable_23_ 23 23.23"
	lexer := NewLexer(data)

	tests := []struct {
		name       string
		token      string
		identifier int
	}{
		{
			name:       "Variable definition",
			token:      "var",
			identifier: TokenVar,
		},
		{
			name:       "Identifier definition",
			token:      "_variable_23_",
			identifier: TokenIdentifier,
		},
		{
			name:       "Integer number definition",
			token:      "23",
			identifier: TokenNumber,
		},
		{
			name:       "Float number definition",
			token:      "23.23",
			identifier: TokenNumber,
		},
		{
			name:       "EOF",
			token:      "",
			identifier: TokenEOF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.token, lexer.Token)
			assert.Equal(t, tt.identifier, lexer.TokenIdentifier)

			lexer.NextToken()
		})
	}

}
