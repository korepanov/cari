package command

/*
grammar

expression = term | expression "+" term | expression "-" term
term = factor | term "*" factor | term "/" factor
factor = number | "(" expression ")"
*/

import (
	"fmt"

	"github.com/korepanov/cari/internal/grammar"
	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myast"
	"github.com/korepanov/cari/internal/myerrors"
)

func factor(in []lexemes.Token) (ast myast.Node, err error) {
	if len(in) == 0 {
		return myast.Node{}, myerrors.ErrNoFactor
	}

	if len(in) == 1 {
		if in[0].T == lexemes.NumberLexeme {
			var node myast.Node
			node.Value = in[0]
			return node, nil
		}
	}

	if in[0].T == lexemes.Delimiter && in[0].Lex == "(" &&
		in[len(in)-1].T == lexemes.Delimiter && in[len(in)-1].Lex == ")" {
		return parse(in[1 : len(in)-1])
	}

	return myast.Node{}, myerrors.ErrNoFactor
}

func term(in []lexemes.Token) (ast myast.Node, err error) {
	factorAst, err := factor(in)
	i := 1

	for err != nil && i < len(in) {
		factorAst, err = factor(in[i:])
		i++
	}
	i--

	if err != nil {
		return
	}

	if i > 0 {
		token := in[i-1]
		if token.T == lexemes.Operator && (token.Lex == "*" || token.Lex == "/") {
			t := in[:i-1]

			var termAst myast.Node

			ast.Value = token
			termAst, err = term(t)

			if err != nil {
				return
			}

			ast.Children = append(ast.Children, &termAst)
			ast.Children = append(ast.Children, &factorAst)

			return ast, nil
		}

	}

	if i == 0 {
		return factorAst, nil
	}

	return ast, myerrors.ErrNoTerm
}

func expr(in []lexemes.Token) (ast myast.Node, err error) {
	termAst, err := term(in)
	i := 1

	for err != nil && i < len(in) {
		termAst, err = term(in[i:])
		i++
	}
	i--

	if err != nil {
		return
	}

	if i > 0 {
		token := in[i-1]
		if token.T == lexemes.Operator && (token.Lex == "+" || token.Lex == "-") {
			t := in[:i-1]

			var exprAst myast.Node

			ast.Value = token
			exprAst, err = expr(t)

			if err != nil {
				return
			}

			ast.Children = append(ast.Children, &exprAst)
			ast.Children = append(ast.Children, &termAst)

			return ast, nil
		}

	}

	if i == 0 {
		return termAst, nil
	}

	return ast, myerrors.ErrNoExpr
}

func parse(in []lexemes.Token) (ast myast.Node, err error) {
	ast, err = expr(in)
	return
}

func (c *Command) Parse() (err error) {
	c.Ast, err = parse(c.Tokens)

	if err != nil {
		return fmt.Errorf("%s : %s", myerrors.ErrParse, err)
	}

	return nil
}

type grammarNode struct {
	T     grammar.GrammarType
	Token []lexemes.Token
}
