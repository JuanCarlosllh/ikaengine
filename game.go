package ikaengine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/juancarlosllh/ikaengine/nodes"
)

type WindowSize struct {
	Width  int
	Height int
}

type InternalResolution struct {
	Width  int
	Height int
}

type DebugConfig struct {
	DisplayFps bool
}

type GameConfig struct {
	WindowTitle        string
	WindowSize         WindowSize
	InternalResolution InternalResolution
	DebugConfig        DebugConfig
}

type Game struct {
	*GameConfig
	Children []nodes.LiveNode
}

func (g *Game) Update() error {
	for _, children := range g.Children {
		children.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.DebugConfig.DisplayFps {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%.0f", ebiten.ActualTPS()))
	}
	for _, children := range g.Children {
		children.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.InternalResolution.Width, g.InternalResolution.Height
}
