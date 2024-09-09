package command

import (
	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myast"
)

type Command struct {
	Input     string
	subinput  string // to copy Input and work with it inside package
	Tokens    []lexemes.Token
	subtokens []lexemes.Token
	Ast       myast.Node
}
