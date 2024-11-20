package dsl

import "testing"

func TestParseReplaceCommandValidSyntax(t *testing.T) {
	parser := &Parser{
		tokens: []Token{
			{Value: "find"},
			{Value: ":"},
			{Value: "replace"},
			{Value: "="},
			{Value: "file.txt"},
		},
	}

	find, replace, file, err := parser.ParseReplaceCommand()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if find != "find" || replace != "replace" || file != "file.txt" {
		t.Fatalf("unexpected result: find=%s, replace=%s, file=%s", find, replace, file)
	}
}

func TestParseReplaceCommandFewerTokens(t *testing.T) {
	parser := &Parser{
		tokens: []Token{
			{Value: "find"},
			{Value: ":"},
			{Value: "replace"},
		},
	}

	_, _, _, err := parser.ParseReplaceCommand()

	if err == nil {
		t.Fatal("expected an error, got none")
	}

	expectedError := "invalid command syntax"
	if err.Error() != expectedError {
		t.Fatalf("expected error %q, got %q", expectedError, err.Error())
	}
}
