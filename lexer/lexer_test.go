package lexer

import (
	"strings"
	"testing"

	"github.com/schafer14/monkey/token"
)

type lexTest struct {
	name   string
	phrase string
	tokens []token.Token
}

var tests = []lexTest{
	lexTest{
		name:   "Basic token test",
		phrase: `=+(){},;`,
		tokens: []token.Token{
			token.Token{
				Type: token.ASSIGN,
			},
			token.Token{
				Type: token.PLUS,
			},
			token.Token{
				Type: token.LPAREN,
			},
			token.Token{
				Type: token.RPAREN,
			},
			token.Token{
				Type: token.LBRACE,
			},
			token.Token{
				Type: token.RBRACE,
			},
			token.Token{
				Type: token.COMMA,
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.EOF,
			},
		},
	},
	lexTest{
		name: "Basic monkey syntax",
		phrase: `let five = 5;
		let ten = 10;
		
		let add = fn(x, y) {
			x + y;
		};
		
		let result = add(five, ten);`,
		tokens: []token.Token{
			token.Token{
				Type: token.LET,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "five",
			},
			token.Token{
				Type: token.ASSIGN,
			},
			token.Token{
				Type:    token.INT,
				Literal: "5",
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.LET,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "ten",
			},
			token.Token{
				Type: token.ASSIGN,
			},
			token.Token{
				Type:    token.INT,
				Literal: "10",
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.LET,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "add",
			},
			token.Token{
				Type: token.ASSIGN,
			},
			token.Token{
				Type: token.FUNCTION,
			},
			token.Token{
				Type: token.LPAREN,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "x",
			},
			token.Token{
				Type: token.COMMA,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "y",
			},
			token.Token{
				Type: token.RPAREN,
			},
			token.Token{
				Type: token.LBRACE,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "x",
			},
			token.Token{
				Type: token.PLUS,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "y",
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.RBRACE,
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.LET,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "result",
			},
			token.Token{
				Type: token.ASSIGN,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "add",
			},
			token.Token{
				Type: token.LPAREN,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "five",
			},
			token.Token{
				Type: token.COMMA,
			},
			token.Token{
				Type:    token.IDENT,
				Literal: "ten",
			},
			token.Token{
				Type: token.RPAREN,
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.EOF,
			},
		},
	},
	lexTest{
		name: "Monkey extended syntax",
		phrase: `!-/*5;
		5 < 10 > 5;`,
		tokens: []token.Token{
			token.Token{
				Type: token.BANG,
			},
			token.Token{
				Type: token.MINUS,
			},
			token.Token{
				Type: token.SLASH,
			},
			token.Token{
				Type: token.ASTERISK,
			},
			token.Token{
				Type:    token.INT,
				Literal: "5",
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type:    token.INT,
				Literal: "5",
			},
			token.Token{
				Type: token.LT,
			},
			token.Token{
				Type:    token.INT,
				Literal: "10",
			},
			token.Token{
				Type: token.GT,
			},
			token.Token{
				Type:    token.INT,
				Literal: "5",
			},
			token.Token{
				Type: token.SEMICOLON,
			},
		},
	},
	lexTest{
		name: "Monkey if else return syntax",
		phrase: `if (5 < 10) {
			return true;
		} else {
			return false;
		}`,
		tokens: []token.Token{
			token.Token{
				Type: token.IF,
			},
			token.Token{
				Type: token.LPAREN,
			},
			token.Token{
				Type:    token.INT,
				Literal: "5",
			},
			token.Token{
				Type: token.LT,
			},
			token.Token{
				Type:    token.INT,
				Literal: "10",
			},
			token.Token{
				Type: token.RPAREN,
			},
			token.Token{
				Type: token.LBRACE,
			},
			token.Token{
				Type: token.RETURN,
			},
			token.Token{
				Type: token.TRUE,
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.RBRACE,
			},
			token.Token{
				Type: token.ELSE,
			},
			token.Token{
				Type: token.LBRACE,
			},
			token.Token{
				Type: token.RETURN,
			},
			token.Token{
				Type: token.FALSE,
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type: token.RBRACE,
			},
		},
	},
	lexTest{
		name: "Monkey comparisions",
		phrase: `10 == 10;
		10 != 9;`,
		tokens: []token.Token{
			token.Token{
				Type:    token.INT,
				Literal: "10",
			},
			token.Token{
				Type: token.EQ,
			},
			token.Token{
				Type:    token.INT,
				Literal: "10",
			},
			token.Token{
				Type: token.SEMICOLON,
			},
			token.Token{
				Type:    token.INT,
				Literal: "10",
			},
			token.Token{
				Type: token.NOT_EQ,
			},
			token.Token{
				Type:    token.INT,
				Literal: "9",
			},
			token.Token{
				Type: token.SEMICOLON,
			},
		},
	},
}

func TestNext(t *testing.T) {
	for _, tt := range tests {
		lex := New(strings.NewReader(tt.phrase))

		for i, expected := range tt.tokens {
			got := lex.Next()

			if got.Type != expected.Type {
				t.Errorf(
					"Failed %v token %v: expected type `%v` got `%v`",
					tt.name,
					i,
					expected.Type,
					got.Type,
				)
			}

			if got.Literal != expected.Literal {
				t.Errorf(
					"Failed %v token %v: expected literal `%v` got `%v`",
					tt.name,
					i,
					expected.Literal,
					got.Literal,
				)
			}
		}
	}
}
