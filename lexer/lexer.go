package lexer

import (
	"io"
	"text/scanner"

	"github.com/schafer14/monkey/token"
)

var EOF = rune(-1)

type Lexer struct {
	s *scanner.Scanner
}

func New(r io.Reader) Lexer {
	s := scanner.Scanner{}
	return Lexer{s: s.Init(r)}
}

func (l Lexer) Next() token.Token {
	l.eatWhitespace()
	pos := l.s.Pos()

	char := l.s.Peek()

	switch char {
	case '!':
		pos := l.s.Pos()
		l.s.Next()
		if l.s.Peek() == '=' {
			l.s.Next()
			return token.Token{Length: 2, Pos: pos, Type: token.NOT_EQ}
		}
		return token.Token{Length: 1, Pos: pos, Type: token.BANG}
	case '-':
		return l.singleToken(token.MINUS)
	case '/':
		return l.singleToken(token.SLASH)
	case '*':
		return l.singleToken(token.ASTERISK)
	case '<':
		return l.singleToken(token.LT)
	case '>':
		return l.singleToken(token.GT)
	case '=':
		pos := l.s.Pos()
		l.s.Next()
		if l.s.Peek() == '=' {
			l.s.Next()
			return token.Token{Length: 2, Pos: pos, Type: token.EQ}
		}
		return token.Token{Length: 1, Pos: pos, Type: token.ASSIGN}
	case '+':
		return l.singleToken(token.PLUS)
	case '(':
		return l.singleToken(token.LPAREN)
	case ')':
		return l.singleToken(token.RPAREN)
	case '{':
		return l.singleToken(token.LBRACE)
	case '}':
		return l.singleToken(token.RBRACE)
	case ',':
		return l.singleToken(token.COMMA)
	case ';':
		return l.singleToken(token.SEMICOLON)
	case EOF:
		return l.singleToken(token.EOF)
	default:
		if isLetter(char) {
			ident, length := l.readIdentifier()
			tokType := token.LookupTokenType(ident)

			if tokType != token.IDENT {
				ident = ""
			}

			return token.Token{
				Length:  length,
				Pos:     pos,
				Type:    tokType,
				Literal: ident,
			}
		}
		if isDigit(char) {
			digit, tokType, length := l.readDigit()

			return token.Token{
				Length:  length,
				Pos:     pos,
				Type:    tokType,
				Literal: digit,
			}
		}
	}

	return token.Token{Type: token.ILLEGAL, Literal: string(char), Pos: pos}
}

func (l *Lexer) readIdentifier() (str string, length uint) {
	for {
		char := l.s.Peek()

		if !isLetter(char) {
			break
		}

		l.s.Next()
		length += 1
		str = str + string(char)
	}

	return str, length
}

func (l *Lexer) readDigit() (str string, tokType token.TokenType, length uint) {
	for {
		char := l.s.Peek()

		if !isDigit(char) {
			break
		}

		l.s.Next()
		length += 1
		str = str + string(char)
	}

	return str, token.INT, length
}

func (l *Lexer) singleToken(t token.TokenType) token.Token {
	pos := l.s.Pos()
	l.s.Next()
	return token.Token{Length: 1, Pos: pos, Type: t}
}

func isLetter(char rune) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) eatWhitespace() {
	for {
		char := l.s.Peek()

		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			l.s.Next()
		} else {
			break
		}
	}
}
