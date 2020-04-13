package tree

// Record is a struct that contains a root id and parent id
type Record struct {
	ID     int
	Parent int
}

// Node is a struct that
type Node struct {
	ID       int
	Children []*Node
}

// Build is a function that builds a tree
func Build(records []Record) (*Node, error) {
	
}
