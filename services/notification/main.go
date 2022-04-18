package main

import (
	"log"
	"sync"

	"github.com/jonashiltl/sessions-backend/packages/events"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/notification/internal"
	"github.com/jonashiltl/sessions-backend/services/notification/internal/config"
	"github.com/jonashiltl/sessions-backend/services/notification/internal/handler"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Println("No .env file found")
	}

	opts := []nats.Option{nats.Name("Notification Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	st := stream.New(nc)

	smtp, err := internal.Connect(c)
	if err != nil {
		log.Fatalln(err)
	}

	s := handler.NewServer(smtp)

	var wg sync.WaitGroup
	wg.Add(1)

	go st.SubscribeByEvent("notification.email.friend.requested", events.FriendRequested{}, s.FriendRequested)
	go st.SubscribeByEvent("notification.email.verify", events.Registered{}, s.Registered)

	// this will wait until the wg counter is at 0
	wg.Wait()
}
