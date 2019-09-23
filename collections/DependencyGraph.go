package collections

type ItemDependencyGraphNodeSet struct {
	items map[*ItemDependencyGraphNode]bool
}

func (s *ItemDependencyGraphNodeSet) Add(value *ItemDependencyGraphNode) bool {
	if s.Contains(value) {
		return false
	}
	s.items[value] = true
	return true
}

func (s *ItemDependencyGraphNodeSet) Remove(value *ItemDependencyGraphNode) bool {
	if s.Contains(value) {
		delete(s.items, value)
		return true
	}
	return false
}

func (s *ItemDependencyGraphNodeSet) Contains(value *ItemDependencyGraphNode) bool {
	_, ok := s.items[value]
	return ok
}

type ItemDependencyGraphNode struct {
	data        ItemValue
	antecedents *ItemDependencyGraphNodeSet
	dependents  *ItemDependencyGraphNodeSet
}

type ItemDependencyGraph struct {
	nodes map[ItemValue]*ItemDependencyGraphNode
}

func NewItemDependencyGraph() *ItemDependencyGraph {
	return &ItemDependencyGraph{
		nodes: make(map[ItemValue]*ItemDependencyGraphNode),
	}
}

func (g *ItemDependencyGraph) Add(value ItemValue) {
	g.add(value)
}

func (g *ItemDependencyGraph) AddDependency(antecedent, dependent ItemValue) {
	fromNode := g.add(antecedent)
	toNode := g.add(dependent)

	fromNode.dependents.Add(fromNode)
	toNode.antecedents.Add(toNode)
}

func (g *ItemDependencyGraph) add(value ItemValue) *ItemDependencyGraphNode {
	var n *ItemDependencyGraphNode
	var ok bool
	if n, ok = g.nodes[value]; !ok {
		n = &ItemDependencyGraphNode{
			data:        value,
			antecedents: &ItemDependencyGraphNodeSet{},
			dependents: &ItemDependencyGraphNodeSet{},
		}
		g.nodes[value] = n
	}
	return n
}