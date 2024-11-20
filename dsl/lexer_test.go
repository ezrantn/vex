package dsl

import "testing"

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
