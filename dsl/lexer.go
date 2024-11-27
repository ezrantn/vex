package dsl

import "strings"

type TokenType string

const (
	TOKEN_EOF                TokenType = "EOF"                //
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
	l.skipWhitespace()

	if l.pos >= len(l.input) {
		return Token{Type: TOKEN_EOF, Value: ""}
	}

	// Handle ':='
	if l.pos+1 < len(l.input) && l.input[l.pos:l.pos+2] == ":=" {
		l.pos += 2
		return Token{Type: TOKEN_SPECIAL_ASSIGNMENT, Value: ":="}
	}

	if l.input[l.pos] == '"' {
		return l.readQuotedString()
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

	// Handle unquoted arguments
	return l.readUnquotedArgument()
}

func (l *Lexer) skipWhitespace() {
	for l.pos < len(l.input) && (l.input[l.pos] == ' ' || l.input[l.pos] == '\t') {
		l.pos++
	}
}

func (l *Lexer) readQuotedString() Token {
	start := l.pos + 1
	l.pos++
	escaped := false
	for l.pos < len(l.input) {
		if l.input[l.pos] == '"' && !escaped {
			break
		}
		if l.input[l.pos] == '\\' {
			escaped = !escaped
		} else {
			escaped = false
		}
		l.pos++
	}

	// Check if the closing quote is found
	if l.pos < len(l.input) && l.input[l.pos] == '"' {
		value := l.input[start:l.pos]
		l.pos++
		return Token{Type: TOKEN_ARG, Value: value}
	}

	// If no closing quote, return the rest of the input
	return Token{Type: TOKEN_ARG, Value: l.input[start:]}
}

func (l *Lexer) readUnquotedArgument() Token {
	start := l.pos
	for l.pos < len(l.input) && !strings.ContainsRune(` :=`, rune(l.input[l.pos])) {
		if l.input[l.pos] == ':' || l.input[l.pos] == '=' {
			break
		}
		l.pos++
	}
	return Token{Type: TOKEN_ARG, Value: l.input[start:l.pos]}
}
