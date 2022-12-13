package ikaengine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/juancarlosllh/ikaengine/nodes"
	"log"
)

type World struct {
	GameConfig *GameConfig
	Children   []nodes.InstantiableNode
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

	var game = &Game{
		GameConfig: config,
		Children:   w.Children,
	}

	for _, children := range w.Children {
		children.Init()
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
