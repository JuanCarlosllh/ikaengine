package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type InstantiableNode interface {
	Init()
	Draw(screen *ebiten.Image)
	Update()
	DrawChild(screen *ebiten.Image)
	UpdateChild()
	AddChild(child InstantiableNode)
	GetParent() *InstantiableNode
	SetParent(child *InstantiableNode)
}

type Node struct {
	Children []InstantiableNode
	Parent   *InstantiableNode
}

func (n *Node) UpdateChild() {
	for _, child := range n.Children {
		child.Update()
		child.UpdateChild()
	}
}

func (n *Node) DrawChild(screen *ebiten.Image) {
	for _, child := range n.Children {
		child.Draw(screen)
		child.DrawChild(screen)
	}
}

func (n *Node) AddChild(child InstantiableNode) {
	child.Init()
	n.Children = append(n.Children, child)
}

func (n *Node) GetParent() *InstantiableNode {
	return n.Parent
}

func (n *Node) SetParent(parent *InstantiableNode) {
	n.Parent = parent
}
