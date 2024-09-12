package myast

import (
	"fmt"

	"github.com/korepanov/cari/internal/lexemes"
	"github.com/korepanov/cari/internal/myerrors"
	"github.com/korepanov/cari/pkg/mytools"
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

func NewNode(v lexemes.Token) Node {
	var n Node
	n.Value = v
	return n
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

/*
The MustAppendNode appends node v to the node with id=parentId of the ast a.
Panics if fails to find the node with id=parentId in the ast a.
*/
func (a *Ast) MustAppendNode(parentId int, v *Node) (id int) {
	node, err := a.Node(parentId)
	if err != nil {
		panic(fmt.Errorf("%s : %s, id: %d", myerrors.ErrAppendNode, err, parentId))
	}

	node.Children = append(node.Children, v)
	v.Parent = node
	v.id = a.nextId
	a.nextId++

	return v.id
}

/*
The MustAppend appends the root of the ast v to the node with parentId of the ast a.
Panics is fails to find the node with id=parentId in the ast a.
*/
func (a *Ast) MustAppend(parentId int, v *Ast) {
	var parents []*Node
	parents = append(parents, v.Root)
	id := parentId

	for _, parent := range parents {
		for _, child := range parent.Children {
			if child.Parent.id == v.Root.id {
				a.MustAppendNode(id, child)
			} else {
				child.id = a.nextId
				a.nextId++
			}
			parents = append(parents, child)
		}
	}
}

func (n *Node) MyId() int {
	return n.id
}

func (a *Ast) Node(id int) (*Node, error) {
	return a.Root.findNodeById(id)
}

/*
The findNodeById returns *Node by it's id scanning the nodes from n as root.
Returns ErrNoNode if there is no node with such id.
*/
func (n *Node) findNodeById(id int) (res *Node, err error) {
	if n.id == id {
		return n, nil
	}
	for _, child := range n.Children {
		res, err = child.findNodeById(id)
		if err == nil {
			return
		}
	}

	return n, myerrors.ErrNoNode
}

/*
The printInLevel prints the lexeme of the node to the terminal with format of the level in the ast.
Returns branchLevels where ├─ or │ was printed.
Wants prevBranchLevels which is the branchLevels from the previos call.
Set prevBranchLevels = []int{} if there was no previous call.
*/
func (n *Node) printInLevel(level int, prevBranchLevels []int) (branchLevels []int) {

	for i := 0; i < level; i++ {
		if mytools.Contains(prevBranchLevels, i) {
			fmt.Print("│  ")
			branchLevels = append(branchLevels, i)
		} else {
			fmt.Print("   ")
		}
	}

	if level != 0 {
		if n.Parent == nil || n.Parent.Children[len(n.Parent.Children)-1].id == n.id {
			fmt.Print("└─ ")
		} else {
			fmt.Print("├─ ")
			branchLevels = append(branchLevels, level)
		}
	}

	fmt.Println(n.Value.Lex)

	for _, child := range n.Children {
		branchLevels = child.printInLevel(level+1, branchLevels)
	}

	return
}

// The Print prints the ast in the terminal.
func (a *Ast) Print() {
	a.Root.printInLevel(0, []int{})
}
