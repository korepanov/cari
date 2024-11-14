package program

import (
	"fmt"

	"github.com/korepanov/cari/internal/myast"
	"github.com/korepanov/cari/internal/sysinfo"
)

type asmString string
type operation string

var idToAsm = make(map[int]asmString)

func (p *Program) makeComment() {
	fmt.Printf("# This code was made by %s version %s\n", sysinfo.Name, sysinfo.Version)
}

func (p *Program) makeData() {
	fmt.Print(dataBegin)

	terminalNodes := p.Ast.Root.TerminalNodes()

	for i := 0; i < len(terminalNodes); i++ {
		idToAsm[terminalNodes[i].Id()] = asmString(fmt.Sprintf("(t%d)", i))
		fmt.Printf("t%d:\n.quad %s\n", i, terminalNodes[i].Value.Lex)
	}
}

func (p *Program) makeBss() {
	fmt.Print(bssBegin)

	var maxNonTerminalLen int

	for _, child := range p.Ast.Root.Children {
		nonTerminalNodes := child.NonTerminalNodes()

		for i := 0; i < len(nonTerminalNodes); i++ {
			idToAsm[nonTerminalNodes[i].Id()] = asmString(fmt.Sprintf("(res%d)", i))
		}

		nonTerminalLen := len(nonTerminalNodes)
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

	for _, child := range p.Ast.Root.Children {
		codeOperation(child)
	}

	fmt.Print(textEnd)
}

func codeOperation(n *myast.Node) asmString {

	if len(n.Children) == 0 {
		return idToAsm[n.Id()]
	}

	if len(n.Children) != 2 {
		panic(fmt.Errorf("invalid child number: %d", len(n.Children)))
	}

	leftChild := n.Children[0]
	rightChild := n.Children[1]
	var leftString asmString = idToAsm[leftChild.Id()]
	var rightString asmString = idToAsm[rightChild.Id()]

	if len(leftChild.Children) > 0 {
		leftString = codeOperation(leftChild)
	}
	if len(rightChild.Children) > 0 {
		rightString = codeOperation(rightChild)
	}

	codeFuncs := []func(operation, asmString, asmString, asmString) bool{
		plus,
		minus,
		mul,
		div,
	}

	var ok bool

	for _, f := range codeFuncs {
		ok = f(operation(n.Value.Lex), idToAsm[n.Id()], leftString, rightString)
		if ok {
			break
		}
	}

	if !ok {
		panic(fmt.Errorf("no such operation: %s", n.Value.Lex))
	}

	return idToAsm[n.Id()]
}

func plus(op operation, res asmString, a asmString, b asmString) bool {
	if op != "+" {
		return false
	}
	fmt.Println(res, "=", a, op, b)
	return true
}

func minus(op operation, res asmString, a asmString, b asmString) bool {
	if op != "-" {
		return false
	}
	fmt.Println(res, "=", a, op, b)
	return true
}

func mul(op operation, res asmString, a asmString, b asmString) bool {
	if op != "*" {
		return false
	}
	fmt.Println(res, "=", a, op, b)
	return true
}

func div(op operation, res asmString, a asmString, b asmString) bool {
	if op != "/" {
		return false
	}
	fmt.Println(res, "=", a, op, b)
	return true
}
