package signals

type Observer interface {
	OnNotify(KeyPressSignal)
}
