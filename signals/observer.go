package signals

type Observer[T interface{}] interface {
	OnNotify(e Signal[T])
	GetConnectedSignals() map[string]func(signal Signal[T])
	Connect(signal string, callback func(signal Signal[T]))
	Disconnect(signal string)
}

type ObserverComponent[T any] struct {
	connectedSignals map[string]func(signal Signal[T])
}

func (o *ObserverComponent[T]) OnNotify(e Signal[T]) {
	o.connectedSignals[e.Type](e)
}

func (o *ObserverComponent[T]) GetConnectedSignals() map[string]func(signal Signal[T]) {
	return o.connectedSignals
}

func (o *ObserverComponent[T]) Connect(signal string, callback func(signal Signal[T])) {
	o.connectedSignals[signal] = callback
}

func (o *ObserverComponent[T]) Disconnect(signal string) {
	o.connectedSignals[signal] = nil
}

func NewObserver[T any]() Observer[T] {
	return &ObserverComponent[T]{
		connectedSignals: make(map[string]func(signal Signal[T])),
	}
}
