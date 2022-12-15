package ikaengine

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/juancarlosllh/ikaengine/nodes"
	"log"
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

type World struct {
	*GameConfig
	Children []nodes.LiveNode
}

func (w *World) Init() {
	var config = w.GameConfig
	if config.WindowSize.Width <= 0 {
		config.WindowSize.Width = 640
	}
	if config.WindowSize.Height <= 0 {
		config.WindowSize.Height = 480
	}
	if config.InternalResolution.Width <= 0 {
		config.InternalResolution.Width = 320
	}
	if config.InternalResolution.Height <= 0 {
		config.InternalResolution.Height = 240
	}

	ebiten.SetWindowSize(config.WindowSize.Width, config.WindowSize.Height)
	ebiten.SetWindowTitle(config.WindowTitle)

	for _, children := range w.Children {
		children.GetNode().Init()
		children.Init()
	}

	if err := ebiten.RunGame(w); err != nil {
		log.Fatal(err)
	}
}

func (w *World) Update() error {
	for _, children := range w.Children {
		children.GetNode().Update()
	}
	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	if w.DebugConfig.DisplayFps {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("%.0f", ebiten.ActualTPS()))
	}
	for _, children := range w.Children {
		children.GetNode().Draw(screen)
	}
}

func (w *World) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w.InternalResolution.Width, w.InternalResolution.Height
}
