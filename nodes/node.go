package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reflect"
	"strings"
)

type Node interface {
	// Life Cycel
	Init()
	Update()
	Draw(screen *ebiten.Image)
	Input()

	// NodeComponent internal live cycle
	RootInit()
	RootUpdate()
	RootDraw(screen *ebiten.Image)
	RootInput()

	// NodeComponent Properties
	GetName() string

	// Structure
	GetParent() Node
	GetChildren() []Node
	AddChild(child Node)
	GetNode() *NodeComponent

	setParent(parent Node)
	setNodeRoot(root Node)
	setName(name string)
}

type NodeComponent struct {
	Node

	Type string
	Name string
	ID   uint

	Children []Node
	Parent   Node
	NodeRoot Node
}

// Life cycle
func (n *NodeComponent) Init()                     {}
func (n *NodeComponent) Update()                   {}
func (n *NodeComponent) Draw(screen *ebiten.Image) {}
func (n *NodeComponent) Input()                    {}

func (n *NodeComponent) RootInit() {
	for _, child := range n.Children {
		child.setParent(n.NodeRoot)
		child.GetNode().RootInit()
		child.Init()
	}
}

func (n *NodeComponent) RootUpdate() {
	for _, children := range n.Children {
		children.GetNode().RootUpdate()
		children.Update()
	}
}

func (n *NodeComponent) RootDraw(screen *ebiten.Image) {
	for _, children := range n.Children {
		children.GetNode().RootDraw(screen)
		children.Draw(screen)
	}
}

func (n *NodeComponent) RootInput() {
	for _, children := range n.Children {
		children.GetNode().RootInput()
		children.RootInput()
	}
}

// Structure
func (n *NodeComponent) AddChild(child Node) {
	child.GetNode().RootInit()
	child.Init()
	child.setParent(n.GetNode().NodeRoot)
	n.Children = append(n.Children, child)
}

func (n *NodeComponent) GetName() string {
	return n.Name
}

func (n *NodeComponent) GetParent() Node {
	return n.Parent
}

func (n *NodeComponent) GetChildren() []Node {
	return n.Children
}

func (n *NodeComponent) GetNode() *NodeComponent {
	return n
}

// private
func (n *NodeComponent) setName(name string) {
	n.Name = name
}

func (n *NodeComponent) setParent(parent Node) {
	n.Parent = parent
}

func (n *NodeComponent) setNodeRoot(root Node) {
	n.NodeRoot = root
}

type NewNodeArgs struct {
	Name     string
	Node     Node
	Children []Node
}

// Global functions and utilities
func NewNode(nodeArgs NewNodeArgs) Node {

	t := reflect.TypeOf(nodeArgs.Node)
	structNameTokens := strings.SplitAfter(t.String(), ".")
	nodeType := structNameTokens[len(structNameTokens)-1]

	rootNode := nodeArgs.Node.GetNode()
	rootNode.NodeRoot = nodeArgs.Node
	rootNode.Name = nodeArgs.Name
	rootNode.Type = nodeType

	for _, child := range nodeArgs.Children {
		nodeArgs.Node.AddChild(child)
	}

	return nodeArgs.Node
}
