package grammar

const (
	// 0 - no grammar
	Expr GrammarType = iota + 1
	Term
	Factor
	Number
	Delimiter
)

type GrammarType int
