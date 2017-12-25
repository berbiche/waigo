package lexer

import "waigo/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
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

// skips whitespace until next token
func (l *Lexer) eatWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// NextToken parses the next token in the string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		tok = newToken(token.BANG, l.ch)
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
		tok = newToken(token.PLUS, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '0':
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isValidLetterForIdentifier(l.ch) {
			tok.Literal = l.readIndentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return tok
}

// New returns a new instance of the lexer struct
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// newToken constructs a new Token struct
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

// isValidLetterForIdentifier returns whether the character is allowed in an identifier
func isValidLetterForIdentifier(ch byte) bool {
	switch {
	case 'a' <= ch && ch <= 'z':
	case 'A' <= ch && ch <= 'Z':
	case ch == '_':
	case ch == '$':
	case ch == '!':
	case ch == '?':
		return true
	}
	return false
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
