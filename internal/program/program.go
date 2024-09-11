package program

import (
	"bufio"
	"fmt"
	"os"

	"github.com/korepanov/cari/internal/command"
	"github.com/korepanov/cari/internal/myast"
	"github.com/korepanov/cari/internal/myerrors"
)

type Program struct {
	Input []command.Command
	Ast   myast.Ast
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
	p.Ast = myast.NewAst()

	for i := 0; i < len(p.Input); i++ {
		err := p.Input[i].Parse()
		if err != nil {
			return fmt.Errorf("%s\n%d\t%s", err, i+1, p.Input[i].Input)
		}
		p.Ast.Append(p.Ast.Root.MyId(), &p.Input[i].Ast)
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
