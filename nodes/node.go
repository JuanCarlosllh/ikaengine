package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juancarlosllh/ikaengine/signals"
	"reflect"
	"strings"
)

type LiveNode interface {
	// Life Cycel
	Init()
	Update()
	Draw(screen *ebiten.Image)
	RootInit()
	RootUpdate()
	RootDraw(screen *ebiten.Image)

	// Node Properties
	GetName() string

	// Structure
	GetParent() LiveNode
	GetChildren() []LiveNode
	AddChild(child LiveNode)
	GetNode() *Node

	// Signals
	EmitSignal()

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

// Life cycle
func (n *Node) Init()                     {}
func (n *Node) Update()                   {}
func (n *Node) Draw(screen *ebiten.Image) {}

func (n *Node) RootInit() {
	for _, child := range n.Children {
		child.setParent(n.NodeRoot)
		child.GetNode().RootInit()
		child.Init()
	}
}

func (n *Node) RootUpdate() {
	for _, children := range n.Children {
		children.GetNode().RootUpdate()
		children.Update()
	}
}
func (n *Node) RootDraw(screen *ebiten.Image) {
	for _, children := range n.Children {
		children.GetNode().RootDraw(screen)
		children.Draw(screen)
	}
}

// Structure
func (n *Node) AddChild(child LiveNode) {
	child.GetNode().RootInit()
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

// Signals
func (n *Node) EmitSignal() {
	signals.EventStream.Notify(signals.KeyPressSignal{
		Key:     "a",
		Pressed: true,
	})
}

// private
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

// Global functions and utilities
func NewNode(nodeArgs NewNodeArgs) LiveNode {

	t := reflect.TypeOf(nodeArgs.Node)
	structNameTokens := strings.SplitAfter(t.String(), ".")
	nodeType := structNameTokens[len(structNameTokens)-1]

	rootNode := nodeArgs.Node.GetNode()
	rootNode.NodeRoot = nodeArgs.Node
	rootNode.Name = nodeArgs.Name
	rootNode.Children = nodeArgs.Children
	rootNode.Type = nodeType

	return nodeArgs.Node
}
