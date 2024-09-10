package myast

import (
	"fmt"

	"github.com/korepanov/cari/internal/lexemes"
)

type Node struct {
	Parent   *Node
	Children []*Node
	Value    lexemes.Token
}

func (n Node) printInLevel(level int) {

	for i := 0; i < level; i++ {
		fmt.Print(" ")
	}

	if level != 0 {
		fmt.Print("|")
	}
	//if len(n.Children) > 0 {

	//}
	fmt.Println(n.Value.Lex)

	for _, child := range n.Children {

		child.printInLevel(level + 1)
	}
}

func (n Node) Print() {
	n.printInLevel(0)
}
