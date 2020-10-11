package parser

// Node ...
type Node struct {
	kind  int
	value *int
	op1   *Node
	op2   *Node
	op3   *Node
}

// NewNode ...
func NewNode(kind int, value *int, op1 *Node, op2 *Node, op3 *Node) *Node {
	return &Node{
		kind:  kind,
		value: value,
		op1:   op1,
		op2:   op2,
		op3:   op3,
	}
}

// Kinds
const (
	KindProgram = iota
	KindVar
	KindIf
)
