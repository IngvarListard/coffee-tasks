package main

import (
	"fmt"
	"unicode"
)

const (
	Body = iota
	Object
	Array
	Number
	String
	Key
	Value
	Null
	Lparen
	Rparen
	Lbrace
	Rbrace
	Comma
	Colon

	EOF
)

func NewParser(text string) *Parser {
	p := &Parser{
		text:         []rune(text),
		currentRune:  0,
		currentIndex: -1,
	}
	p.next()
	return p
}

type Parser struct {
	text         []rune
	currentRune  rune
	currentIndex int
}

func (p *Parser) next() {
	p.currentIndex++
	if p.currentIndex > len(p.text) {
		p.currentRune = EOF
	}
}

func (p *Parser) skipWhitespace() {
	for unicode.IsSpace(p.currentRune) {
		p.next()
	}
}

func (p *Parser) ReadBody() (result interface{}, err error) {

	switch p.currentRune {
	case '[':
		p.next()
		p.skipWhitespace()
		result, err = p.ReadValue()
	case '{':
		p.next()
		p.skipWhitespace()
		result, err = p.ReadObject()
	default:
		return nil, fmt.Errorf("unexpected rune occurrance %s", string(p.currentRune))
	}
	return
}

func (p *Parser) ReadValue() (result interface{}, err error) {
	switch v := p.currentRune; {
	case v == '[':
		p.next()
		p.skipWhitespace()
		result, err = p.ReadValue()
	case v == '{':
		p.next()
		p.skipWhitespace()
		result, err = p.ReadObject()
	case v == '"':
		p.next()
		p.skipWhitespace()
		result, err = p.ReadString()
	case unicode.IsDigit(v):
		p.next()
		p.skipWhitespace()
		result, err = p.ReadNumber()
	case unicode.IsSymbol(v):
		p.next()
		p.skipWhitespace()
		result, err = p.ReadKeyWord()
	default:
		return nil, fmt.Errorf("unexpected char while parsing array: %s", string(v))
	}
	p.next()
	p.skipWhitespace()
}
