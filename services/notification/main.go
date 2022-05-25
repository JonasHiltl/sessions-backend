package main

import (
	"log"
	"sync"

	"github.com/jonashiltl/sessions-backend/packages/events"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/notification/config"
	"github.com/jonashiltl/sessions-backend/services/notification/handler"
	"github.com/jonashiltl/sessions-backend/services/notification/mail"
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

	smtp, err := mail.Connect(c)
	if err != nil {
		log.Fatalln(err)
	}

	s := handler.NewServer(smtp)

	var wg sync.WaitGroup
	wg.Add(3)

	go st.SubscribeByEvent("notification.email.friend.requested", events.FriendRequested{}, s.FriendRequested)
	go st.SubscribeByEvent("notification.email.verify", events.Registered{}, s.Registered)

	go st.SubscribeByEvent("notification.push.party.created", events.PartyCreated{}, s.PartyCreated)

	// this will wait until the wg counter is at 0
	wg.Wait()
}
