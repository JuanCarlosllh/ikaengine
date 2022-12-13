package nodes

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type ImageNode struct {
	Node2D
	Path string
	img  *ebiten.Image
}

func (n *ImageNode) Init() {
	fmt.Println("INIT")
	img, _, err := ebitenutil.NewImageFromFile(n.Path)
	if err != nil {
		log.Fatal(err)
	}
	n.img = img
	n.getPosition()
}

func (n *ImageNode) Update() {
}

func (n *ImageNode) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(n.Position.X, n.Position.Y)
	screen.DrawImage(n.img, op)
}
