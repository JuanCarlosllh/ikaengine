package ikaengine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/juancarlosllh/ikaengine/signals"
)

type InputManager struct {
	JustPressedKeys  []ebiten.Key
	JustReleasedKeys []ebiten.Key
}

func (im *InputManager) Update() {
	Input.JustPressedKeys = inpututil.AppendJustPressedKeys(Input.JustPressedKeys[:0])
	Input.JustReleasedKeys = inpututil.AppendJustReleasedKeys(Input.JustPressedKeys[:0])

	for _, key := range Input.JustPressedKeys {
		eventKeySignal := signals.NewSignal[signals.InputEventKey]("InputEventKey", signals.InputEventKey{
			Keycode: key,
			Pressed: true,
		})
		signals.InputNotifier.Notify(eventKeySignal)
	}

	for _, key := range Input.JustReleasedKeys {
		eventKeySignal := signals.NewSignal[signals.InputEventKey]("InputEventKey", signals.InputEventKey{
			Keycode: key,
			Pressed: false,
		})
		signals.InputNotifier.Notify(eventKeySignal)
	}

}

var Input = &InputManager{}
