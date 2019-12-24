package eventbus

import (
	"context"
	"sync"
)

var (
	eventBus       *EventBus
	once           sync.Once
	channelQLength = 100
	topicQLength   = 100
)

func init() {
	eventBus = &EventBus{
		subscribes: make(map[string]*channel),
		stopC:      make(chan struct{}),
		topicC:     make(chan string, topicQLength),
		publishPool: sync.Pool{
			New: func() interface{} {
				return &eventPublisher{}
			},
		},
	}
	eventBus.broadcastLoop()
}

func PubSub() *EventBus {
	return eventBus
}

// 管道组
type channel struct {
	pubC      chan interface{}
	subCGroup []chan interface{}
}

// 事件总线
type EventBus struct {
	subscribes map[string]*channel

	topicC      chan string
	stopC       chan struct{}
	publishPool sync.Pool
}

// 将发布者发送的事件消息广播至其订阅者
func (eb *EventBus) broadcastLoop() {
	go func() {
		for {
			select {
			case <-eb.stopC:
				return
			case topic := <-eb.topicC:
				channel, found := eb.subscribes[topic]
				if !found {
					continue
				}

				for event := range channel.pubC {
					for _, subC := range channel.subCGroup {
						subC <- event
					}
				}
			}
		}
	}()
}

// 发布者注册topic事件
func (eb *EventBus) Pub(topicName string, queueLength ...int) Publisher {
	if len(queueLength) > 0 {
		channelQLength = queueLength[0]
	}

	c, found := eb.subscribes[topicName]
	if !found {
		pubC := make(chan interface{}, channelQLength)
		c = &channel{pubC: pubC}
		eb.subscribes[topicName] = c
	}

	if c.pubC == nil {
		c.pubC = make(chan interface{}, channelQLength)
	}

	p := eb.publishPool.Get().(*eventPublisher)
	defer eb.publishPool.Put(p)
	p.ch = c.pubC

	eb.topicC <- topicName

	return p
}

// 订阅 没有就创建主题
func (eb *EventBus) Sub(ctx context.Context, topicName string, do func(event interface{})) {
	subC := make(chan interface{}, channelQLength)
	c, found := eb.subscribes[topicName]
	if !found {
		c = &channel{}
		c.subCGroup = []chan interface{}{subC}
		eb.subscribes[topicName] = c
	} else {
		c.subCGroup = append(c.subCGroup, subC)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-eb.stopC:
				return
			case event := <-subC:
				do(event)
			}
		}
	}()
}

func (eb *EventBus) Close() {
	eb.stopC <- struct{}{}
}
