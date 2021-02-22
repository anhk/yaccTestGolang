package calc

import (
	"strconv"
	"strings"
	"text/scanner"
)

type Token struct {
	Type int
	Str  string
}

type Calc struct {
	scanner.Scanner
	value float64
}

func (c *Calc) Error(s string) {
	panic("implement me")
}

func (c *Calc) Lex(lval *yySymType) int {
	r, lit := c.Scan(), c.TokenText()
	var token Token
	token.Str = lit

	switch r {
	case scanner.EOF:
		return 0
	case scanner.Int:
		i, _ := strconv.Atoi(lit)
		lval.num = float64(i)
		token.Type = scanner.Float
	case scanner.Float:
		lval.num, _ = strconv.ParseFloat(lit, 64)
		token.Type = scanner.Float
	default:
		token.Type = int(r)
	}

	if token.Type == scanner.Float {
		return NUMBER
	}
	return token.Type
}

func Parse(code string) float64 {
	s := &Calc{}
	s.Init(strings.NewReader(code))
	yyParse(s)
	return s.value
}
