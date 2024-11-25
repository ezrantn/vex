package dsl

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewParserWithValidInput(t *testing.T) {
	input := "valid input"
	parser := NewParser(input)

	if parser == nil {
		t.Errorf("expected parser to be initialized, got nil")
	}

	if len(parser.tokens) == 0 {
		t.Errorf("expected tokens to be populated, got empty slice")
	}
}

func TestNewParserWithEmptyInput(t *testing.T) {
	input := ""
	parser := NewParser(input)

	if parser == nil {
		t.Errorf("Expected parser to be initialized, got nil")
	}

	if len(parser.tokens) != 1 || parser.tokens[0].Type != TOKEN_EOF {
		t.Errorf("Expected tokens to contain only EOF token, got %v", parser.tokens)
	}
}

func TestExpectReturnsTokenOnMatch(t *testing.T) {
	parser := &Parser{
		tokens: []Token{
			{Type: "IDENTIFIER", Value: "x"},
		},
		pos: 0,
	}
	expectedType := TokenType("IDENTIFIER")
	token, err := parser.expect(expectedType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token.Type != expectedType {
		t.Errorf("expected token type %s, got %s", expectedType, token.Type)
	}
}

func TestExpectHandlesUnexpectedEndOfInput(t *testing.T) {
	parser := &Parser{
		tokens: []Token{},
		pos:    0,
	}
	expectedType := TokenType("IDENTIFIER")
	_, err := parser.expect(expectedType)
	if err == nil {
		t.Fatal("expected an error, got none")
	}
	expectedError := "unexpected end of input"
	if !strings.Contains(err.Error(), expectedError) {
		t.Errorf("expected error to contain %q, got %v", expectedError, err)
	}
}

func TestParseReplaceCommandValidSyntax(t *testing.T) {
	p := &Parser{
		expectFunc: func(tokenType TokenType) (Token, error) {
			switch tokenType {
			case TOKEN_ARG:
				return Token{Value: "value"}, nil
			case TOKEN_COLON, TOKEN_EQUAL:
				return Token{}, nil
			default:
				return Token{}, fmt.Errorf("unexpected token type")
			}
		},
	}

	find, replace, file, err := p.ParseReplaceCommand()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if find != "value" || replace != "value" || file != "value" {
		t.Fatalf("expected 'value', 'value', 'value', got %v, %v, %v", find, replace, file)
	}
}

func TestParseReplaceCommandMissingFindArgument(t *testing.T) {
	p := &Parser{
		expectFunc: func(tokenType TokenType) (Token, error) {
			if tokenType == TOKEN_ARG {
				return Token{}, fmt.Errorf("missing 'find' argument")
			}
			return Token{}, nil
		},
	}

	_, _, _, err := p.ParseReplaceCommand()

	if err == nil || err.Error() != "failed to parse 'find' argument: missing 'find' argument" {
		t.Fatalf("expected error for missing 'find' argument, got %v", err)
	}
}

func TestParseFileCommandValidInput(t *testing.T) {
	parser := &Parser{
		tokens: []Token{
			{Value: ":="},
			{Value: "file_name"},
		},
	}
	result, err := parser.ParseFileCommand()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	expected := "file_name"
	if result != expected {
		t.Errorf("expected %s, got %s", expected, result)
	}
}

func TestParseFileCommandInsufficientTokens(t *testing.T) {
	parser := &Parser{
		tokens: []Token{
			{Value: ":="},
		},
	}
	_, err := parser.ParseFileCommand()
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	expectedError := "syntax error: input is too short. Expected ':=file_name', but got incomplete input"
	if err.Error() != expectedError {
		t.Errorf("expected error %s, got %s", expectedError, err.Error())
	}
}

func TestParseFilterCommandValidInput(t *testing.T) {
	parser := &Parser{
		expectFunc: func(tokenType TokenType) (Token, error) {
			switch tokenType {
			case TOKEN_ARG:
				return Token{Value: "test"}, nil
			case TOKEN_EQUAL:
				return Token{}, nil
			default:
				return Token{}, fmt.Errorf("unexpected token type")
			}
		},
	}

	word, file, err := parser.ParseFilterCommand()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if word != "test" || file != "test" {
		t.Fatalf("expected word and file to be 'test', got word: %s, file: %s", word, file)
	}
}

func TestParseFilterCommandMissingWord(t *testing.T) {
	parser := &Parser{
		expectFunc: func(tokenType TokenType) (Token, error) {
			if tokenType == TOKEN_ARG {
				return Token{}, fmt.Errorf("missing word argument")
			}
			return Token{}, nil
		},
	}

	_, _, err := parser.ParseFilterCommand()

	if err == nil || err.Error() != "failed to parse 'word' argument: missing word argument" {
		t.Fatalf("expected error for missing word argument, got %v", err)
	}
}
