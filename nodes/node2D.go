package nodes

import (
	"github.com/juancarlosllh/ikaengine/math"
)

type Node2DInterface interface {
	GetPosition() math.Vector2
	SetPosition(position math.Vector2)
}

type Node2D struct {
	Node2DInterface
	Position math.Vector2
}

func (n *Node2D) GetPosition() math.Vector2 {
	return n.Position
}

func (n *Node2D) SetPosition(position math.Vector2) {
	n.Position = position
}
