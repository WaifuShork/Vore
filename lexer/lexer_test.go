package lexer

import (
	"testing"

	"vore/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;
	`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Assignment, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
	
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Assignment, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
	
		{token.Let, "let"},
		{token.Identifier, "add"},
		{token.Assignment, "="},
		{token.Function, "fn"},
		{token.LParen, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Identifier, "x"},
		{token.Plus, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.RBrace, "}"},
		{token.Semicolon, ";"},
	
		{token.Let, "let"},
		{token.Identifier, "result"},
		{token.Assignment, "="},
		{token.Identifier, "add"},
		{token.LParen, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RParen, ")"},
		{token.Semicolon, ";"},
	
		// !-/*5;
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		
		// 5 < 10 > 5;
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.GT, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},

		// if (5 < 10) {
		// 		return true;
		// } else {
		// 		return false;
		// }
		{token.If, "if"},
		{token.LParen, "("},
		{token.Int, "5"},
		{token.LT, "<"},
		{token.Int, "10"},
		{token.RParen, ")"},
		{token.LBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.RBrace, "}"},
		{token.Else, "else"},
		{token.LBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.RBrace, "}"},

		// 10 == 10;
		{token.Int, "10"},
		{token.Equal, "=="},
		{token.Int, "10"},
		{token.Semicolon, ";"},

		// 10 != 9;
		{token.Int, "10"},
		{token.NotEqual, "!="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
		
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range tests {
		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - TokenType wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}