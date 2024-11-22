package dsl

import "strings"

type TokenType string

const (
	TOKEN_EOF                TokenType = "EOF"				  //
 	TOKEN_SPECIAL_ASSIGNMENT TokenType = "SPECIAL_ASSIGNMENT" // :=
	TOKEN_COLON              TokenType = "COLON"              // :
	TOKEN_EQUAL              TokenType = "EQUAL"              // =
	TOKEN_ARG                TokenType = "ARG"
)

type Token struct {
	Type  TokenType
	Value string
}

type Lexer struct {
	input string
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		pos:   0,
	}
}

func (l *Lexer) NextToken() Token {
	// Skip whitespace
	for l.pos < len(l.input) && l.input[l.pos] == ' ' {
		l.pos++
	}

	if l.pos >= len(l.input) {
		return Token{Type: TOKEN_EOF, Value: ""}
	}

	// Handle ':='
	if l.pos+1 < len(l.input) && l.input[l.pos:l.pos+2] == ":=" {
		l.pos += 2
		return Token{Type: TOKEN_SPECIAL_ASSIGNMENT, Value: ":="}
	}

	// Handle ':' and '='
	if l.input[l.pos] == ':' {
		l.pos++
		return Token{Type: TOKEN_COLON, Value: ":"}
	}
	if l.input[l.pos] == '=' {
		l.pos++
		return Token{Type: TOKEN_EQUAL, Value: "="}
	}

	// Handle quoted strings
	if l.input[l.pos] == '"' {
		start := l.pos + 1
		l.pos++
		for l.pos < len(l.input) && l.input[l.pos] != '"' {
			l.pos++
		}
		if l.pos < len(l.input) && l.input[l.pos] == '"' {
			l.pos++ // Consume closing quote
		}
		return Token{Type: TOKEN_ARG, Value: l.input[start : l.pos-1]}
	}

	// Handle unquoted arguments
	start := l.pos
	for l.pos < len(l.input) && !strings.ContainsRune(` :=`, rune(l.input[l.pos])) {
		l.pos++
	}
	return Token{Type: TOKEN_ARG, Value: strings.TrimSpace(l.input[start:l.pos])}
}
