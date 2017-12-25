package lexer

import (
	"fmt"
	"waigo/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

// New returns a new instance of the lexer struct
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken parses the next token in the string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isValidLetterForIdentifier(l.ch) {
			tok.Literal = l.readIndentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok = token.Token{Type: token.INT, Literal: l.readNumber()}
			return tok
		}
		fmt.Printf("%q\n", l.ch)
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

// returns a slice of the input containing the identifier
func (l *Lexer) readIndentifier() string {
	position := l.position
	for isValidLetterForIdentifier(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// returns a slice of the input containing the number
func (l *Lexer) readNumber() string {
	initialPosition := l.position

	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[initialPosition:l.position]
}

// returns the next char without advancing the read position in the input
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// skips whitespace until next token
func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// newToken constructs a new Token struct
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isValidLetterForIdentifier returns whether the character is allowed in an identifier
func isValidLetterForIdentifier(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		'A' <= ch && ch <= 'Z' ||
		ch == '_' ||
		ch == '$' ||
		ch == '!' ||
		ch == '?'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
