package signals

type Signal[T any] struct {
	Type string
	Data T
}

type Notifier[T interface{}] interface {
	Register(Observer[T])
	Unregister(Observer[T])
	Notify(e Signal[T])
}

type NotifierComponent[T any] struct {
	observers map[Observer[T]]struct{}
}

func (n *NotifierComponent[T]) Register(o Observer[T]) {
	n.observers[o] = struct{}{}
}

func (n *NotifierComponent[T]) Unregister(o Observer[T]) {
	delete(n.observers, o)
}

func (n *NotifierComponent[T]) Notify(e Signal[T]) {
	for o := range n.observers {
		if o != nil && o.GetConnectedSignals()[e.Type] != nil {
			o.OnNotify(e)
		}
	}
}

func NewNotifier[T any]() Notifier[T] {
	return &NotifierComponent[T]{
		observers: make(map[Observer[T]]struct{}),
	}
}

func NewSignal[T any](signalType string, data T) Signal[T] {
	return Signal[T]{
		Type: signalType,
		Data: data,
	}
}
