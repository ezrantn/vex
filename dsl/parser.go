package dsl

import (
	"fmt"
)

type Parser struct {
	tokens []Token
	pos    int
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
	if p.pos >= len(p.tokens) {
		return Token{}, fmt.Errorf("unexpected end of input")
	}
	token := p.tokens[p.pos]
	if token.Type != expectedType {
		return Token{}, fmt.Errorf("expected %s, got %s", expectedType, token.Type)
	}
	p.pos++
	return token, nil
}

func (p *Parser) ParseReplaceCommand() (find, replace, file string, err error) {
	findToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", "", err
	}
	find = findToken.Value

	_, err = p.expect(TOKEN_COLON)
	if err != nil {
		return "", "", "", err
	}

	replaceToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", "", err
	}
	replace = replaceToken.Value

	_, err = p.expect(TOKEN_EQUAL)
	if err != nil {
		return "", "", "", err
	}

	fileToken, err := p.expect(TOKEN_ARG)
	if err != nil {
		return "", "", "", err
	}
	file = fileToken.Value

	return find, replace, file, nil
}

func (p *Parser) ParseFileCommand() (string, error) {
	if len(p.tokens) < 2 {
		return "", fmt.Errorf("invalid load command syntax")
	}

	// Token 0: := operator
	if p.tokens[0].Value != ":=" {
		return "", fmt.Errorf("expected ':=' for file assignment")
	}

	// Token 1: File name
	file := p.tokens[1].Value
	return file, nil
}
