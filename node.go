package mapCreator

import "github.com/ygto/goui/entity"

type Node struct {
	neighbours []*Node
	value      float64
	nodeType   int
	name       string
	entity     *entity.Entity
}

func NewNode(name string, t int) *Node {
	n := Node{}
	n.name = name
	n.nodeType = t
	n.neighbours = make([]*Node, 0, 0)

	return &n
}
func (node *Node) SetEntity(e *entity.Entity) {
	node.entity = e
}
func (node *Node) GetEntity() *entity.Entity {
	return node.entity
}
func (node *Node) GetName() string {
	return node.name
}
func (node *Node) GetValue() float64 {
	return node.value
}
func (node *Node) SetValue(val float64) {
	node.value = val
}
func (node *Node) GetNodeType() int {
	return node.nodeType
}
func (node *Node) GetNeighbours() []*Node {
	return node.neighbours
}
func (node *Node) addNeighbours(neighbours ... *Node) *Node {

	//check old neighbours to protect multiplexing
	for _, neighbour := range neighbours {
		for _, n := range node.neighbours {
			if n.name == neighbour.name {
				return node
			}
		}
		node.neighbours = append(node.neighbours, neighbour)

		// two way neighbourhood
		neighbour.addNeighbours(node)
	}

	return node
}
