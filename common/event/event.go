package event

import (
	"sync"
)

// Callback ...
type Callback func(eventArgs interface{}, subscriberArgs interface{})

// EventType ...
type EventType int

type request struct {
	callback Callback
	args     interface{}
	once     bool
}

// EventBus ...
type EventBus struct {
	sync.Mutex
	requests map[EventType]map[string]*request
}

func (eb *EventBus) subscribe(t EventType, c Callback, a interface{}, id string) {
	eb.Lock()
	defer eb.Unlock()
	if _, ok := eb.requests[t]; !ok {
		eb.requests[t] = make(map[string]*request)
	}
	eb.requests[t][id] = &request{callback: c, args: a, once: false}
}

func (eb *EventBus) subscribeOnce(t EventType, c Callback, a interface{}, id string) {
	eb.Lock()
	defer eb.Unlock()
	if _, ok := eb.requests[t]; !ok {
		eb.requests[t] = make(map[string]*request)
	}
	eb.requests[t][id] = &request{callback: c, args: a, once: true}
}

func (eb *EventBus) unSubscribe(t EventType, id string) {
	eb.Lock()
	defer eb.Unlock()
	if _, ok := eb.requests[t]; !ok {
		return
	}
	delete(eb.requests[t], id)
}

// TODO Return an array of error in channel
func (eb *EventBus) rise(t EventType, args interface{}) {
	eb.Lock()
	defer eb.Unlock()
	toDelete := []string{}
	if _, ok := eb.requests[t]; !ok {
		return
	}
	for k, v := range eb.requests[t] {
		go v.callback(args, v.args)
		if v.once {
			toDelete = append(toDelete, k)
		}
	}
	for _, id := range toDelete {
		delete(eb.requests[t], id)
	}
}

var eb EventBus

func init() {
	eb.requests = make(map[EventType]map[string]*request)
}

// Subscribe ...
func Subscribe(t EventType, c Callback, args interface{}, id string) {
	eb.subscribe(t, c, args, id)
}

// SubscribeOnce ...
func SubscribeOnce(t EventType, c Callback, args interface{}, id string) {
	eb.subscribeOnce(t, c, args, id)
}

// UnSubscribe ...
func UnSubscribe(t EventType, id string) {
	eb.unSubscribe(t, id)
}

// Rise ...
func Rise(t EventType, args interface{}) {
	eb.rise(t, args)
}
