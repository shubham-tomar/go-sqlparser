package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType string

const (
	EOF        TokenType = "EOF"
	ILLEGAL    TokenType = "ILLEGAL"
	Identifier TokenType = "Identifier"

	IDENT     TokenType = "IDENT" // Table name, column name, etc.
	KEYWORD   TokenType = "KEYWORD"
	DATATYPE  TokenType = "DATATYPE"
	NUMBER    TokenType = "NUMBER"
	STRING    TokenType = "STRING"
	LPAREN    TokenType = "("
	RPAREN    TokenType = ")"
	LBRACE    TokenType = "{"
	RBRACE    TokenType = "}"
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"
	DOT       TokenType = "."
	PLUS      TokenType = "+"
	MINUS     TokenType = "-"
	ASTERISK  TokenType = "*"
	SLASH     TokenType = "/"
	COLON     TokenType = ":"
)

type Token struct {
	Type    TokenType
	Literal string
}

// Lexer struct holds the input SQL string
type Lexer struct {
	input        string
	position     int  // Current character position
	readPosition int  // Next character position
	ch           byte // Current character
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
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

func (l *Lexer) NextToken() (Token, error) {
	l.skipWhitespace()

	switch l.ch {
	case '(':
		l.readChar()
		return Token{Type: LPAREN, Literal: "("}, nil
	case ')':
		l.readChar()
		return Token{Type: RPAREN, Literal: ")"}, nil
	case ',':
		l.readChar()
		return Token{Type: COMMA, Literal: ","}, nil
	case ';':
		l.readChar()
		return Token{Type: SEMICOLON, Literal: ";"}, nil
	case '.':
		l.readChar()
		return Token{Type: DOT, Literal: "."}, nil
	case 0:
		return Token{Type: EOF, Literal: ""}, nil
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tokenType := lookupKeyword(literal)

			return Token{Type: tokenType, Literal: literal}, nil
		} else if isDigit(l.ch) {
			return Token{Type: IDENT, Literal: l.readNumber()}, nil
		} else {
			return Token{Type: ILLEGAL, Literal: string(l.ch)}, fmt.Errorf("unexpected character: '%c'", l.ch)
		}
	}
}

// // Helper: Convert character into token
// func newToken(tokenType TokenType, ch byte) Token {
// 	return Token{Type: tokenType, Literal: string(ch)}
// }

// readIdentifier reads SQL keywords or identifiers
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// readNumber reads numeric literals
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace ignores spaces and newlines
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Helper: Check if character is a letter (A-Z, a-z, _)
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// Helper: Check if character is a digit (0-9)
func isDigit(ch byte) bool {
	return unicode.IsDigit(rune(ch))
}

// Keyword Lookup Table
var keywords = map[string]TokenType{
	"CREATE":    "KEYWORD",
	"TABLE":     "KEYWORD",
	"NAMESPACE": "KEYWORD",
	"INT":       "DATATYPE",
	"STRING":    "DATATYPE",
}

// lookupKeyword checks if a word is a SQL keyword
func lookupKeyword(word string) TokenType {
	if tok, ok := keywords[strings.ToUpper(word)]; ok {
		return tok
	}
	return IDENT
}
