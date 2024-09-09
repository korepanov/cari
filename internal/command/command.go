package command

import (
	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myast"
	"github.com/korepanov/cari/internal/mytypes"
)

type Command struct {
	Input    string
	subinput string // to copy Input and work with it inside package
	Tokens   []lexemes.Token
	Ast      myast.Node
}

type Variable struct {
	T     mytypes.UserType
	Name  lexemes.Lexeme
	Value interface{}
}

type VariablesTable struct {
	variables []Variable
}

func (vt *VariablesTable) Append(v []Variable) {
	vt.variables = append(vt.variables, v...)
}

// returns function to get variable from the end of VariablesTable
// function returns Variable and true if ok
// function returns empty Variable and false if no variables
func (vt *VariablesTable) Iterator() func() (Variable, bool) {
	i := len(vt.variables)

	return func() (Variable, bool) {
		i--
		if i >= 0 {
			return vt.variables[i], true
		}
		return Variable{}, false
	}
}
