package lexer

import (
	"strings"
)

type TokenType string

const (
	BinaryOperator TokenType = "BinaryOperator"
	Identifier     TokenType = "Identifier"
)

/*
 * Token Structure
 */
type Token struct {
	Value string
	Type  TokenType
}

func (t Token) Val() string {
	return t.Value
}

func (t Token) Typ() TokenType {
	return t.Type
}

func token(value string, Typ TokenType) Token {
	return Token{Value: value, Type: Typ}
}

func tokenize(sourcecode string) []Token {
	src := strings.Split(sourcecode, "")
	Tokens := []Token{}

	for len(src) > 0 {
		if src[0] == "+" {
			Tokens = append(Tokens, token(src[0], BinaryOperator))
			src = src[1:]
		}
	}

	return Tokens
}

/*
func main() {
	fmt.Println(tokenize("+"))
}
*/
