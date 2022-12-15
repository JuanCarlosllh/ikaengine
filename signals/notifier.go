package signals

type Notifier interface {
	Register(Observer)
	Unregister(Observer)
	Notify(KeyPressSignal)
}

type GlobalNotifier struct {
	Observers map[Observer]struct{}
}

func (n *GlobalNotifier) Register(o Observer) {
	n.Observers[o] = struct{}{}
}

func (n *GlobalNotifier) Unregister(o Observer) {
	delete(n.Observers, o)
}

func (n *GlobalNotifier) Notify(e KeyPressSignal) {
	for o := range n.Observers {
		o.OnNotify(e)
	}
}

var EventStream = GlobalNotifier{
	Observers: map[Observer]struct{}{},
}
