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
	drawPosition := n.Position
	if parent, ok := n.GetParent().(Node2DInterface); ok {
		drawPosition.X = parent.GetPosition().X + n.Position.X
		drawPosition.Y = parent.GetPosition().Y + n.Position.Y
	}
	op.GeoM.Translate(drawPosition.X, drawPosition.Y)
	screen.DrawImage(n.img, op)
}
