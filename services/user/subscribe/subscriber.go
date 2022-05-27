package subscribe

import (
	"log"
	"sync"

	"github.com/jonashiltl/sessions-backend/packages/events"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/user/service"
)

type subscriber struct {
	stream stream.Stream
	ps     service.ProfileService
}

type Subscriber interface {
	IncrementFriendCount(*events.FriendAccepted)
	DecrementFriendCount(*events.FriendRemoved)

	Start()
}

func NewSubscriber(stream stream.Stream, ps service.ProfileService) Subscriber {
	return &subscriber{stream: stream, ps: ps}
}

func (s subscriber) Start() {
	var wg sync.WaitGroup
	wg.Add(2)

	go s.stream.SubscribeToEvent("userservice.increment.friendcount", events.FriendAccepted{}, s.IncrementFriendCount)
	go s.stream.SubscribeToEvent("userservice.decrement.friendcount", events.FriendRemoved{}, s.DecrementFriendCount)

	log.Println("Subscribed for events")

	wg.Wait()
}

func (s subscriber) IncrementFriendCount(e *events.FriendAccepted) {

	log.Printf("%v", e)

}

func (s subscriber) DecrementFriendCount(e *events.FriendRemoved) {

	log.Printf("%v", e)

}
