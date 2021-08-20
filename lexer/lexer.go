package lexer

import "vore/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	char         byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // ensure the lexer is fully operational before returning it
	return l
}


func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0 // if we're reached end, set our current ASCII character to NUL, signifying we're done reading
	} else {
		l.char = l.input[l.readPosition] // set the current character to the next character, effectively advancing our lexer
	}

	l.position = l.readPosition // actually advance the lexer to it's next position
	l.readPosition += 1         // make sure this is advanced so it's always 1 ahead
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{Type: token.Equal, Literal: string(char) + string(l.char)};	
		} else {
			tok = newToken(token.Assignment, l.char)
		}
	case '!':
		if l.peekChar() == '=' {
			char := l.char
			l.readChar()
			tok = token.Token{Type: token.NotEqual, Literal: string(char) + string(l.char)}
		} else {
			tok = newToken(token.Bang, l.char)
		}
	case ';':
		tok = newToken(token.Semicolon, l.char)
	case '(':
		tok = newToken(token.LParen, l.char)
	case ')':
		tok = newToken(token.RParen, l.char)
	case ',':
		tok = newToken(token.Comma, l.char)
	case '+':
		tok = newToken(token.Plus, l.char)
	case '{':
		tok = newToken(token.LBrace, l.char)
	case '}':
		tok = newToken(token.RBrace, l.char)
	case '-':
		tok = newToken(token.Minus, l.char)
	case '/':
		tok = newToken(token.Slash, l.char)
	case '*':
		tok = newToken(token.Asterisk, l.char)
	case '<':
		tok = newToken(token.LT, l.char)
	case '>':
		tok = newToken(token.GT, l.char)	
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.char) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.char) {
			tok.Type = token.Int
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.Illegal, l.char)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

// We just want to view the value, not actually increment our current position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for (isDigit(l.char)) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	// just read forever while it's whitespace since we don't want the whitespace tokens 
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 
		   'A' <= char && char <= 'Z' || 
		   char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}