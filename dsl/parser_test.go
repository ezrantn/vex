package dsl

import (
	"fmt"
	"reflect"
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

func TestParseReplaceCommandSingleArguments(t *testing.T) {
	p := &Parser{}
	p.tokens = []Token{
		{Type: TOKEN_ARG, Value: "findValue"},
		{Type: TOKEN_COLON, Value: ":"},
		{Type: TOKEN_ARG, Value: "replaceValue"},
		{Type: TOKEN_EQUAL, Value: "="},
		{Type: TOKEN_ARG, Value: "fileValue"},
	}

	findList, replaceList, file, err := p.ParseReplaceCommand()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedFindList := []string{"findValue"}
	expectedReplaceList := []string{"replaceValue"}
	expectedFile := "fileValue"

	if !reflect.DeepEqual(findList, expectedFindList) {
		t.Errorf("expected findList %v, got %v", expectedFindList, findList)
	}

	if !reflect.DeepEqual(replaceList, expectedReplaceList) {
		t.Errorf("expected replaceList %v, got %v", expectedReplaceList, replaceList)
	}

	if file != expectedFile {
		t.Errorf("expected file %v, got %v", expectedFile, file)
	}
}

func TestParseReplaceCommandMultipleArguments(t *testing.T) {
    p := &Parser{}
    p.tokens = []Token{
        {Type: TOKEN_ARG, Value: "find1,find2"},
        {Type: TOKEN_COLON, Value: ":"},
        {Type: TOKEN_ARG, Value: "replace1,replace2"},
        {Type: TOKEN_EQUAL, Value: "="},
        {Type: TOKEN_ARG, Value: "fileValue"},
    }
    
    findList, replaceList, file, err := p.ParseReplaceCommand()
    
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }
    
    expectedFindList := []string{"find1", "find2"}
    expectedReplaceList := []string{"replace1", "replace2"}
    expectedFile := "fileValue"
    
    if !reflect.DeepEqual(findList, expectedFindList) {
        t.Errorf("expected findList %v, got %v", expectedFindList, findList)
    }
    
    if !reflect.DeepEqual(replaceList, expectedReplaceList) {
        t.Errorf("expected replaceList %v, got %v", expectedReplaceList, replaceList)
    }
    
    if file != expectedFile {
        t.Errorf("expected file %v, got %v", expectedFile, file)
    }
}

func TestParseReplaceCommandMissingColon(t *testing.T) {
    p := &Parser{}
    p.tokens = []Token{
        {Type: TOKEN_ARG, Value: "findValue"},
        // Missing TOKEN_COLON here
        {Type: TOKEN_ARG, Value: "replaceValue"},
        {Type: TOKEN_EQUAL, Value: "="},
        {Type: TOKEN_ARG, Value: "fileValue"},
    }
    
    _, _, _, err := p.ParseReplaceCommand()
    
    if err == nil || !strings.Contains(err.Error(), "missing ':' after 'find' argument") {
        t.Fatalf("expected error about missing ':', got %v", err)
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
