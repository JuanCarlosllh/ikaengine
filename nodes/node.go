package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type LiveNode interface {
	Init()
	Update()
	Draw(screen *ebiten.Image)
	GetName() string

	GetParent() LiveNode
	setParent(parent LiveNode)
	GetChildren() []LiveNode
	AddChild(child LiveNode)

	setNodeRoot(root LiveNode)
	setName(name string)
}

type Node struct {
	LiveNode
	Type     string
	Name     string
	Children []LiveNode
	Parent   LiveNode
	NodeRoot LiveNode
}

func (n *Node) Init() {}
func (n *Node) Update() {
	for _, children := range n.Children {
		children.Update()
	}
}
func (n *Node) Draw(screen *ebiten.Image) {
	for _, children := range n.Children {
		children.Draw(screen)
	}
}
func (n *Node) GetName() string {
	return n.Name
}
func (n *Node) setName(name string) {
	n.Name = name
}

func (n *Node) setParent(parent LiveNode) {
	n.Parent = parent
}
func (n *Node) GetParent() LiveNode {
	return n.Parent
}
func (n *Node) GerChildren() []LiveNode {
	return n.Children
}
func (n *Node) AddChild(child LiveNode) {
	child.setNodeRoot(child)
	child.Init()
	child.setParent(n.NodeRoot)
	n.Children = append(n.Children, child)
}

func (n *Node) setNodeRoot(root LiveNode) {
	n.NodeRoot = root
}
