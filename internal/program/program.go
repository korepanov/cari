package program

import (
	"bufio"
	"fmt"
	"os"

	"github.com/korepanov/cari/internal/command"
	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myast"
	"github.com/korepanov/cari/internal/myerrors"
)

type Program struct {
	Input []command.Command
	Ast   myast.Node
}

func (p *Program) ReadProgram() error {

	s := bufio.NewScanner(os.Stdin)

	for p.nextCommand(s) {
	}

	err := p.lexicalAnalyze()
	if err != nil {
		return fmt.Errorf("%s : %s", myerrors.ErrRead, err)
	}

	err = p.parse()

	if err != nil {
		return fmt.Errorf("%s : %s", myerrors.ErrRead, err)
	}

	return nil
}

func (p *Program) lexicalAnalyze() error {
	for i := 0; i < len(p.Input); i++ {
		err := p.Input[i].LexicalAnalyze()
		if err != nil {
			return fmt.Errorf("%s\n%d\t%s", err, i+1, p.Input[i].Input)
		}
	}
	return nil
}

func (p *Program) parse() error {
	var root lexemes.Token
	root.Lex = "start"
	root.T = lexemes.StartLexeme
	p.Ast.Value = root

	for i, command := range p.Input {
		err := command.Parse()
		if err != nil {
			return fmt.Errorf("%s\n%d\t%s", err, i+1, command.Input)
		}
	}

	return nil
}

func (p *Program) WriteProgram() {
	for _, command := range p.Input {
		for _, token := range command.Tokens {
			fmt.Print(token.Lex + " ; ")
		}
		fmt.Println()
	}
}

// reads next command in the command input
func (p *Program) nextCommand(s *bufio.Scanner) bool {
	if !s.Scan() {
		return false
	}
	var c command.Command
	c.Input = s.Text()
	p.Input = append(p.Input, c)

	return true
}
