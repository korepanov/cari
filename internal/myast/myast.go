package myast

import "github.com/korepanov/cari/internal/lexemes"

type Node struct {
	Parent   *Node
	Children []*Node
	Value    lexemes.Token
}
