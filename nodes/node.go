package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juancarlosllh/ikaengine/signals"
	"reflect"
	"strings"
)

type LiveNode interface {
	Init()
	Update()
	Draw(screen *ebiten.Image)
	GetName() string

	GetParent() LiveNode
	GetChildren() []LiveNode
	AddChild(child LiveNode)
	GetNode() *Node

	setParent(parent LiveNode)
	setNodeRoot(root LiveNode)
	setName(name string)
}

type Node struct {
	LiveNode
	signals.Observer

	Type     string
	Name     string
	ID       uint
	Children []LiveNode
	Parent   LiveNode
	NodeRoot LiveNode
}

func (n *Node) Init() {
	for _, child := range n.Children {
		child.setNodeRoot(child)
		child.setParent(n.NodeRoot)
		child.GetNode().Init()
		child.Init()
	}
}

func (n *Node) Update() {
	for _, children := range n.Children {
		children.GetNode().Update()
		children.Update()
	}
}
func (n *Node) Draw(screen *ebiten.Image) {
	for _, children := range n.Children {
		children.GetNode().Draw(screen)
		children.Draw(screen)
	}
}

func (n *Node) AddChild(child LiveNode) {
	child.Init()
	n.Children = append(n.Children, child)
}

func (n *Node) GetName() string {
	return n.Name
}

func (n *Node) GetParent() LiveNode {
	return n.Parent
}

func (n *Node) GetChildren() []LiveNode {
	return n.Children
}

func (n *Node) GetNode() *Node {
	return n
}

func (n *Node) setName(name string) {
	n.Name = name
}

func (n *Node) setParent(parent LiveNode) {
	n.Parent = parent
}

func (n *Node) setNodeRoot(root LiveNode) {
	n.NodeRoot = root
}

type NewNodeArgs struct {
	Name     string
	Node     LiveNode
	Children []LiveNode
}

func NewNode(nodeArgs NewNodeArgs) LiveNode {

	t := reflect.TypeOf(nodeArgs.Node)
	structNameTokens := strings.SplitAfter(t.String(), ".")
	nodeType := structNameTokens[len(structNameTokens)-1]

	rootNode := nodeArgs.Node.GetNode()
	rootNode.Name = nodeArgs.Name
	rootNode.Children = nodeArgs.Children
	rootNode.GetNode().Type = nodeType
	return nodeArgs.Node
}
