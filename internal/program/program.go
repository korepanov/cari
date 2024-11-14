package program

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/korepanov/cari/internal/command"
	"github.com/korepanov/cari/internal/myast"
	"github.com/korepanov/cari/internal/myerrors"
)

type Program struct {
	Input []command.Command
	Ast   myast.Ast
}

/*
The ReadProgram reads the program from os.Stdin.
*/
func (p *Program) ReadProgram() {
	s := bufio.NewScanner(os.Stdin)

	for p.nextCommand(s) {
	}
}

/*
The AnalyzeProgram makes lexical analysis and parses the program.
*/
func (p *Program) AnalyzeProgram() error {
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
		p.Ast.MustAppend(p.Ast.Root.Id(), &p.Input[i].Ast)
	}

	return nil
}

/*
The WriteProgram writes the program to the os.Stdout.
*/
func (p *Program) WriteProgram() {
	p.makeComment()
	p.makeData()
	p.makeBss()
}

/*
The nextCommand reads next command in the command input.
*/
func (p *Program) nextCommand(s *bufio.Scanner) bool {
	if !s.Scan() {
		return false
	}
	var c command.Command
	c.Input = removeSpaces(s.Text())
	if len(c.Input) > 0 {
		p.Input = append(p.Input, c)
	}

	return true
}

func removeSpaces(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}
