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

func (n Node) Print() {
	fmt.Println("ast")
	/*if len(n.Children) > 0{
		fmt.Println("|")
	}*/
}
