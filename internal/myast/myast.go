package myast

import (
	"fmt"

	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myerrors"
)

type Node struct {
	id       int
	Parent   *Node
	Children []*Node
	Value    lexemes.Token
}

type Ast struct {
	nextId int
	Root   *Node
}

func NewAst() Ast {
	var ast Ast
	var node Node
	ast.Root = &node
	ast.Root.id = ast.nextId
	ast.nextId++

	ast.Root.Value.Lex = "start"
	ast.Root.Value.T = lexemes.StartLexeme

	return ast
}

func (a *Ast) AppendNode(parentId int, v *Node) {
	node, err := a.Node(parentId)
	if node == nil {
		panic(fmt.Errorf("%s, id: %d", myerrors.ErrNoNode, parentId))
	}
	node.Children = append(node.Children, v)
	v.Parent = node
	v.id = a.nextId
	a.nextId++
	fmt.Printf("%v\n", node)
	fmt.Printf("%v\n", a.Root)
}

func (a *Ast) Append(parentId int, v *Ast) {
	for _, child := range v.Root.Children {
		a.AppendNode(parentId, child)
	}
}

/*
	func (a *Ast) Append(parentId int, v *Ast) {
		for i := 0; i < len(v.Root.Children); i++ {
			a.lastNode.Children = append(a.lastNode.Children, v.Root.Children[i])
			v.Root.Children[i].Parent = a.lastNode
			v.Root.Children[i].Id += a.nextId
		}

		a.lastNode = v.Root.Children[len(v.Root.Children)-1]
		a.nextId += len(v.Root.Children)
	}
*/
func (n Node) MyId() int {
	return n.id
}

func (a *Ast) Node(id int) (Node, error) {
	return a.Root.findNodeById(id)
}

func (n Node) findNodeById(id int) (Node, error) {
	if n.id == id {
		return n, nil
	}
	for _, child := range n.Children {
		child.findNodeById(id)
	}

	return n, myerrors.ErrNoNode
}

func (n Node) printInLevel(level int) {

	for i := 0; i < level; i++ {
		fmt.Print("   ")
	}

	if level != 0 {
		if n.Parent == nil || n.Parent.Children[len(n.Parent.Children)-1] == &n {
			fmt.Print("└─ ")
		} else {
			fmt.Print("├─ ")
		}
	}
	//if len(n.Children) > 0 {

	//}
	fmt.Println(n.Value.Lex)

	for _, child := range n.Children {

		child.printInLevel(level + 1)
	}
}

func (a *Ast) Print() {
	a.Root.printInLevel(0)
}
