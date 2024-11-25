package dsl

import "testing"

func TestNewLexerWithValidInput(t *testing.T) {
	input := "valid input"
	lexer := NewLexer(input)

	if lexer.input != input {
		t.Errorf("Expected input to be %s, got %s", input, lexer.input)
	}

	if lexer.pos != 0 {
		t.Errorf("Expected initial position to be 0, got %d", lexer.pos)
	}
}

func TestNewLexerWithEmptyInput(t *testing.T) {
	input := ""
	lexer := NewLexer(input)

	if lexer.input != input {
		t.Errorf("Expected input to be an empty string, got %s", lexer.input)
	}

	if lexer.pos != 0 {
		t.Errorf("Expected initial position to be 0, got %d", lexer.pos)
	}
}

func TestNextTokenReturnsEOFOnEmptyInput(t *testing.T) {
	lexer := Lexer{input: "", pos: 0}
	token := lexer.NextToken()
	if token.Type != TOKEN_EOF {
		t.Errorf("Expected TOKEN_EOF, got %v", token.Type)
	}
}

func TestNextTokenHandlesWhitespaceOnlyInput(t *testing.T) {
	lexer := Lexer{input: "   ", pos: 0}
	token := lexer.NextToken()
	if token.Type != TOKEN_EOF {
		t.Errorf("Expected TOKEN_EOF, got %v", token.Type)
	}
}
