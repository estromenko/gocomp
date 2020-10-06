package lexer

// Tokens
const (
	TokenEOF = iota - 1

	TokenIf
	TokenElse
	TokenVar
	TokenFor
	TokenWhile
	TokenFunction

	TokenNumber
	TokenIdentifier

	TokenLeftBrace
	TokenRightBrace
	TokenLeftPar
	TokenRightPar
	TokenLeftBracket
	TokenRightBracket

	TokenSemicolon
	TokenColon
	TokenDot
	TokenComma
	TokenLess
	TokenMore
	TokenPlus
	TokenMinus
	TokenEqual
	TokenNot
	TokenDivision
	TokenMod

	TokenQuote
	TokenDoubleQuote

	TokenAnd
	TokenOr
)

// Tokens ...
var Tokens = map[string]int{
	"if":    TokenIf,
	"else":  TokenElse,
	"var":   TokenVar,
	"for":   TokenFor,
	"while": TokenWhile,
	"fn":    TokenFunction,

	"{": TokenLeftBrace,
	"}": TokenRightBrace,
	"(": TokenLeftPar,
	")": TokenRightPar,
	"[": TokenLeftBracket,
	"]": TokenRightBracket,

	";": TokenSemicolon,
	":": TokenColon,
	".": TokenDot,
	",": TokenComma,
	"<": TokenLess,
	">": TokenMore,
	"+": TokenPlus,
	"-": TokenMinus,
	"=": TokenEqual,
	"!": TokenNot,
	"/": TokenDivision,
	"%": TokenMod,

	`'`: TokenQuote,
	`"`: TokenDoubleQuote,

	"&": TokenAnd,
	"|": TokenOr,
}
