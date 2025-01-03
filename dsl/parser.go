package dsl

import (
	"fmt"
	"strings"
)

type Parser struct {
	tokens     []Token
	pos        int
	expectFunc func(expectedType TokenType) (Token, error)
}

func NewParser(input string) *Parser {
	lexer := NewLexer(input)
	tokens := []Token{}

	for {
		token := lexer.NextToken()
		tokens = append(tokens, token)
		if token.Type == TOKEN_EOF {
			break
		}
	}

	return &Parser{tokens: tokens}
}

func (p *Parser) expect(expectedType TokenType) (Token, error) {
	if p.expectFunc != nil {
		return p.expectFunc(expectedType)
	}

	if p.pos >= len(p.tokens) {
		return Token{}, fmt.Errorf("unexpected end of input: was expecting '%s' at position %d", expectedType, p.pos)
	}
	token := p.tokens[p.pos]
	if token.Type != expectedType {
		return Token{}, fmt.Errorf("syntax error at position %d, expected '%s' but found '%s' (value: '%s')", p.pos, expectedType, token.Type, token.Value)
	}
	p.pos++
	return token, nil
}

func (p *Parser) ParseReplaceCommand() (findList, replaceList []string, file string, err error) {
	findToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to parse 'find' argument: %v", err)
	}

	nextToken := p.tokens[p.pos]
	if nextToken.Type != TOKEN_COLON && nextToken.Type != TOKEN_EQUAL {
		return nil, nil, "", fmt.Errorf("expected ':' or '=' after find argument")
	}
	p.pos++

	replaceToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to parse 'replace' argument: %v", err)
	}

	if nextToken.Type == TOKEN_COLON {
		_, err = p.expect(TOKEN_EQUAL)
		if err != nil {
			return nil, nil, "", fmt.Errorf("missing '=' after 'replace' argument: %v", err)
		}
	}

	fileToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return nil, nil, "", fmt.Errorf("failed to parse 'file' argument: %v", err)
	}
	file = fileToken.Value

	if strings.Contains(findToken.Value, ",") || strings.Contains(replaceToken.Value, ",") {
		findList = strings.Split(findToken.Value, ",")
		replaceList = strings.Split(replaceToken.Value, ",")
		if len(findList) != len(replaceList) {
			return nil, nil, "", fmt.Errorf("'find' and 'replace' lists must have the same number of elements")
		}
	} else {
		findList = []string{findToken.Value}
		replaceList = []string{replaceToken.Value}
	}

	return findList, replaceList, file, nil
}

func (p *Parser) ParseFileCommand() (string, error) {
	if len(p.tokens) < 2 {
		return "", fmt.Errorf("syntax error: input is too short. Expected ':=file_name', but got incomplete input")
	}

	if p.tokens[0].Value != ":=" {
		return "", fmt.Errorf("syntax error: expected ':=' at the beginning of the command but found '%s'", p.tokens[0].Value)
	}

	file := p.tokens[1].Value
	if file == "" {
		return "", fmt.Errorf("syntax error: file name cannot be empty after ':='. Please provide a valid file name")
	}

	return file, nil
}

func (p *Parser) ParseFilterCommand() (word, file string, err error) {
	wordToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse 'word' argument: %w", err)
	}

	word = wordToken.Value

	_, err = p.expect(TOKEN_EQUAL)
	if err != nil {
		return "", "", fmt.Errorf("failed to find '=' operator: %w", err)
	}

	fileToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse 'file' argument: %w", err)
	}

	file = fileToken.Value

	return word, file, nil
}

func (p *Parser) ParseRegexCommand() (pattern, file string, err error) {
	patternToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse the regex: %w", err)
	}

	pattern = patternToken.Value

	_, err = p.expect(TOKEN_EQUAL)
	if err != nil {
		return "", "", fmt.Errorf("failed to find '=' operator: %w", err)
	}

	fileToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse 'file' argument: %w", err)
	}

	file = fileToken.Value

	return pattern, file, nil
}

func (p *Parser) ParseCountCommand() (word, file string, err error) {
	wordToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse 'count' argument: %w", err)
	}

	word = wordToken.Value

	_, err = p.expect(TOKEN_EQUAL)
	if err != nil {
		return "", "", fmt.Errorf("failed to find '=' operator: %w", err)
	}

	fileToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse 'file' argument: %w", err)
	}

	file = fileToken.Value

	return word, file, nil
}

func (p *Parser) ParseFormatType() (sourceType string, toType string, err error) {
	sourceTypeToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse 'sourceType' argument: %w", err)
	}

	sourceType = sourceTypeToken.Value

	_, err = p.expect(TOKEN_EQUAL)
	if err != nil {
		return "", "", fmt.Errorf("failed to find '=' operator: %w", err)
	}

	toTypeToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", fmt.Errorf("failed to parse 'toType' argument: %w", err)
	}

	toType = toTypeToken.Value

	return sourceType, toType, nil
}