package mapCreator

import (
	"fmt"
	"github.com/ygto/goui"
	"github.com/ygto/goui/entity"
)

const (
	TILE_WALL   = iota
	TILE_ROAD   = iota
	TILE_GOAL   = iota
	TILE_DANGER = iota
)

type Map2D struct {
	x     int
	y     int
	nodes []*Node
}

func NewMap2D(s *Schema, scene *goui.Scene, tileSize int) *Map2D {
	m := new(Map2D)
	m.nodes = make([]*Node, 0, 0)
	m.y = s.GetY()
	m.x = s.GetX()
	for y := 0; y < m.y; y++ {
		for x := 0; x < m.x; x++ {
			node := NewNode(fmt.Sprintf("m-%d:%d", y, x), s.GetSchema(y, x))
			rect := entity.CreateRectangle(float32(x*tileSize), float32(y*tileSize), float32(tileSize), float32(tileSize))
			rect.SetColor(GetColor(node.GetNodeType()))
			scene.AddEntity(rect)
			node.SetEntity(rect)
			m.nodes = append(m.nodes, node)
		}
	}
	for y := 0; y < m.y; y++ {
		for x := 0; x < m.x; x++ {
			index := y*m.x + x
			node := m.nodes[index]
			if node.GetNodeType() == 0 {
				continue
			}
			//left
			leftNode := m.GetNode(y, x-1)
			if leftNode != nil && x != 0 && leftNode.GetNodeType() != 0 {
				node.addNeighbours(leftNode)
			}

			//right
			rightNode := m.GetNode(y, x+1)
			if rightNode != nil && x != m.x-1 && rightNode.GetNodeType() != 0 {
				node.addNeighbours(rightNode)
			}

			//top
			topNode := m.GetNode(y-1, x)
			if topNode != nil && topNode.GetNodeType() != 0 {
				node.addNeighbours(topNode)
			}
			//bottom
			bottomNode := m.GetNode(y+1, x)
			if bottomNode != nil && bottomNode.GetNodeType() != 0 {
				node.addNeighbours(bottomNode)
			}

		}
	}

	return m
}

func (m *Map2D) GetNodes() []*Node {

	return m.nodes
}

func (m *Map2D) GetNode(y int, x int) *Node {
	index := y*m.x + x
	if index < 0 || index >= m.y*m.x {
		return nil
	}
	return m.nodes[index]
}

func (m *Map2D) GetY() int {

	return m.y
}

func (m *Map2D) GetX() int {

	return m.x
}
