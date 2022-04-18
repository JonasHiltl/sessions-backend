package subscribe

import (
	"context"
	"log"
	"sync"

	"github.com/jonashiltl/sessions-backend/packages/events"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/user/internal/service"
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

func (s *subscriber) Start() {
	var wg sync.WaitGroup
	wg.Add(2)

	go s.stream.SubscribeByEvent("userservice.increment.friendcount", events.FriendAccepted{}, s.IncrementFriendCount)
	go s.stream.SubscribeByEvent("userservice.decrement.friendcount", events.FriendRemoved{}, s.DecrementFriendCount)

	log.Println("Subscribed for events")

	wg.Wait()
}

func (s *subscriber) IncrementFriendCount(e *events.FriendAccepted) {
	ctx := context.Background()

	log.Printf("%v", e)

	s.ps.IncrementFriendCount(ctx, e.UserId)
}

func (s *subscriber) DecrementFriendCount(e *events.FriendRemoved) {
	ctx := context.Background()

	log.Printf("%v", e)

	s.ps.DecrementFriendCount(ctx, e.UserId)
}
