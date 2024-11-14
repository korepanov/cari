package program

import (
	"fmt"

	"github.com/korepanov/cari/internal/sysinfo"
)

func (p *Program) makeComment() {
	fmt.Printf("# This code was made by %s version %s\n", sysinfo.Name, sysinfo.Version)
}

func (p *Program) makeData() {
	fmt.Print(dataBegin)

	terminalNodes := p.Ast.Root.TerminalNodes()

	for _, node := range terminalNodes {
		fmt.Printf("t%d:\n.quad %s\n", node.Id(), node.Value.Lex)
	}
}

func (p *Program) makeBss() {
	fmt.Print(bssBegin)

	var maxNonTerminalLen int

	for _, child := range p.Ast.Root.Children {
		nonTerminalLen := len(child.NonTerminalNodes())
		if nonTerminalLen > maxNonTerminalLen {
			maxNonTerminalLen = nonTerminalLen
		}
	}

	for i := 0; i < maxNonTerminalLen; i++ {
		fmt.Printf("res%d:\n.skip 21\n", i)
	}

}

func (p *Program) makeText() {
	fmt.Print(textBegin)
	fmt.Print(textEnd)
}
