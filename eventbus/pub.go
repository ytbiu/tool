package eventbus

type Publisher interface {
	Publish(event interface{})
}

type eventPublisher struct {
	ch chan interface{}
}

func (p *eventPublisher) Publish(event interface{}) {
	p.ch <- event
}
