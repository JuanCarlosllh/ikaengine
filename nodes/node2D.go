package nodes

import (
	"fmt"
	"github.com/juancarlosllh/ikaengine/math"
	"reflect"
)

type Node2D struct {
	Node
	Position math.Vector2
}

func (n *Node2D) getPosition() math.Vector2 {
	t := reflect.TypeOf(n.Parent)
	fmt.Println(t)
	if n.Parent == nil {
		return n.Position
	}
	return math.Vector2{}
}
