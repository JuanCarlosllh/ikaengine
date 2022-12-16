package signals

import "github.com/hajimehoshi/ebiten/v2"

type InputEventKey struct {
	Keycode ebiten.Key
	Pressed bool
}

var InputNotifier = NewNotifier[InputEventKey]()
