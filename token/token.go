// Token structure for goasm65816
package token

type TokenType int

type Token struct {
	Type  TokenType
	Text  string
	Line  int
	Index int
	File  string
}
