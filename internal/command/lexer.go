package command

import (
	"fmt"

	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myerrors"
)

// The lookAhead looks ahead to the next shortest token but does not remove the token from subinput.
func (c *Command) lookAhead() (lexemes.Token, error) {
	var buf string

	for _, ch := range c.subinput {
		buf += string(ch)
		bufType := Dictionary().Find(lexemes.Lexeme(buf))

		if bufType != 0 {
			return lexemes.Token{Lex: lexemes.Lexeme(buf), T: bufType}, nil
		}
	}

	if len(c.subinput) == 0 {
		return lexemes.Token{}, nil
	}

	return lexemes.Token{Lex: lexemes.Lexeme(buf), T: 0}, myerrors.ErrNoToken
}

// The nextToken gets next longest token and removes this token from subinput.
func (c *Command) nextToken(prevToken lexemes.Token) (lexemes.Token, error) {

	var newToken lexemes.Token

loop:
	for aheadToken, err := c.lookAhead(); (err == myerrors.ErrNoToken || !Dictionary().IsStop(aheadToken.T)) &&
		len(aheadToken.Lex) != 0; aheadToken, err = c.lookAhead() {

		buf := string(newToken.Lex)

		for idx, ch := range c.subinput {
			buf += string(ch)
			bufType := Dictionary().Find(lexemes.Lexeme(buf))

			if bufType != 0 {
				newToken.Lex = lexemes.Lexeme(buf)
				newToken.T = bufType
				c.subinput = c.subinput[idx+1:]
				break
			}

			if idx == len(c.subinput)-1 {
				break loop
			}
		}
	}

	if len(newToken.Lex) == 0 {
		newToken, err := c.lookAhead()

		if err != nil {
			return newToken, err
		}

		c.subinput = c.subinput[len(newToken.Lex):]

		if newToken.Lex == "-" && prevToken.T != lexemes.NumberLexeme { // unary minus
			newToken, err = c.nextToken(newToken)
			newToken.Lex = "-" + newToken.Lex
		}

		return newToken, err
	}

	return newToken, nil

}

func (c *Command) LexicalAnalyze() error {

	c.subinput = c.Input
	var newToken lexemes.Token

	for newToken, err := c.nextToken(newToken); len(newToken.Lex) > 0; newToken, err = c.nextToken(newToken) {

		if err != nil {
			return fmt.Errorf("%s : %s", myerrors.ErrLexAnalysis, err)
		}

		c.Tokens = append(c.Tokens, newToken)
	}

	return nil
}
