package program

import (
	"fmt"

	"github.com/korepanov/cari/internal/sysinfo"
)

const dataBegin = `
.data
enter:
.ascii "\n"
.space 1, 0
`

func (p *Program) makeComment() {
	fmt.Printf("# This code was made by %s version %s\n", sysinfo.Name, sysinfo.Version)
}

func (p *Program) makeData() {
	fmt.Print(dataBegin)

	terminalNodes := p.Ast.Root.TerminalNodes()

	for _, node := range terminalNodes {
		fmt.Printf("t%d:\n.quad %s\n", node.MyId(), node.Value.Lex)
	}

	var maxTerminalLen int

	for _, child := range p.Ast.Root.Children {
		terminalLen := len(child.TerminalNodes())
		if terminalLen > maxTerminalLen {
			maxTerminalLen = terminalLen
		}
	}

	fmt.Println(maxTerminalLen)
}
