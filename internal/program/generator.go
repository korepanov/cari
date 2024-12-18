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
		fmt.Printf("res%d:\n.skip 8\n", i)
	}

}

func (p *Program) makeText() {
	fmt.Print(textBegin)

	for _, child := range p.Ast.Root.Children {
		makePrint(codeOperation(child))
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
		add,
		sub,
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

func add(op operation, res asmString, a asmString, b asmString) bool {
	if op != "+" {
		return false
	}
	fmt.Printf("mov %s, %%rax\n", a)
	fmt.Printf("mov %s, %%rbx\n", b)
	fmt.Println("add %rbx, %rax")
	fmt.Printf("mov %%rax, %s\n", res)

	return true
}

func sub(op operation, res asmString, a asmString, b asmString) bool {
	if op != "-" {
		return false
	}
	fmt.Printf("mov %s, %%rax\n", a)
	fmt.Printf("mov %s, %%rbx\n", b)
	fmt.Println("sub %rbx, %rax")
	fmt.Printf("mov %%rax, %s\n", res)

	return true
}

func mul(op operation, res asmString, a asmString, b asmString) bool {
	if op != "*" {
		return false
	}
	fmt.Printf("mov %s, %%rax\n", a)
	fmt.Printf("mov %s, %%rbx\n", b)
	fmt.Println("imul %rbx, %rax")
	fmt.Printf("mov %%rax, %s\n", res)

	return true
}

func div(op operation, res asmString, a asmString, b asmString) bool {
	if op != "/" {
		return false
	}
	fmt.Printf("mov %s, %%rax\n", a)
	fmt.Printf("mov %s, %%rbx\n", b)
	fmt.Println("cqo")
	fmt.Println("idiv %rbx")
	fmt.Printf("mov %%rax, %s\n", res)
	return true
}

func makePrint(res asmString) {
	fmt.Printf("mov %s, %%rax\n", res)
	fmt.Print(printConst)
	fmt.Println()
}
