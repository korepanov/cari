package command

/*
grammar

expression = term | expression "+" term | expression "-" term
term = factor | term "*" factor | term "/" factor
factor = number | "(" expression ")"
*/

import (
	"github.com/korepanov/cari/internal/grammar"
	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myast"
	"github.com/korepanov/cari/internal/myerrors"
)

func factor(in []lexemes.Token) (ast myast.Node, out []lexemes.Token, err error) {
	if len(in) == 1 {
		if in[0].T == lexemes.NumberLexeme {
			var node myast.Node
			node.Value = in[0]
			return node, []lexemes.Token{}, nil
		}
	}

	if in[0].T == lexemes.Delimiter && in[0].Lex == "(" &&
		in[len(in)-1].T == lexemes.Delimiter && in[len(in)-1].Lex == ")" {
		return parse(in)
	}

	return myast.Node{}, []lexemes.Token{}, myerrors.ErrNoFactor
}

func term(in []lexemes.Token) (ast myast.Node, out []lexemes.Token, err error) {
	factorAst, out, err := factor(in)
	i := 1

	for err != nil && i > 0 {
		factorAst, _, err = factor(in[i:])
		if err != nil {
			i++
		}
	}

	if err != nil {
		return
	}

	//f := in[i:]
	if i > 0 {
		token := in[i-1]
		if token.T == lexemes.Operator && (token.Lex == "*" || token.Lex == "/") {

		}

	}

	return
}

func parse(expr []lexemes.Token) (ast myast.Node, out []lexemes.Token, err error) {
	return term(expr)
}

func (c *Command) Parse() error {
	parse(c.Tokens)

	return nil
}

type grammarNode struct {
	T     grammar.GrammarType
	Token []lexemes.Token
}
