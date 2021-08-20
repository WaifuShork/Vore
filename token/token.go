package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

var keyword = map[string]TokenType{
	"if":     If,
	"let":    Let,
	"else":   Else,
	"true":   True,
	"false":  False,
	"return": Return,
	"fn":     Function,
}

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	Identifier = "IDENTIFIER"
	Int        = "INT"

	Assignment = "="
	Plus       = "+"
	Minus      = "-"
	Bang       = "!"
	Asterisk   = "*"
	Slash      = "/"

	LT = "<"
	GT = ">"

	Comma     = ","
	Semicolon = ";"

	LParen = "("
	RParen = ")"
	LBrace = "{"
	RBrace = "}"

	Equal    = "=="
	NotEqual = "!="

	Function = "FUNCTION"
	Let      = "LET"
	True     = "TRUE"
	False    = "FALSE"
	If       = "IF"
	Else     = "ELSE"
	Return   = "RETURN"
)

// attempt to lookup a keyword in the dictionary, or scan an identifier
func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keyword[identifier]; ok {
		return tok
	}

	return Identifier
}