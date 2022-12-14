package nodes

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type ImageNode struct {
	Node
	Node2D
	Path string
	img  *ebiten.Image
}

func (n *ImageNode) Init() {
	img, _, err := ebitenutil.NewImageFromFile(n.Path)
	if err != nil {
		log.Fatal(err)
	}
	n.img = img
}

func (n *ImageNode) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	if parent, ok := n.GetParent().(Node2DInterface); ok {
		op.GeoM.Translate(parent.GetPosition().X+n.Position.X, parent.GetPosition().Y+n.Position.Y)
	} else {
		op.GeoM.Translate(n.Position.X, n.Position.Y)
	}
	screen.DrawImage(n.img, op)
	n.Node.Draw(screen)
}
