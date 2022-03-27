package main

import (
	"log"
	"sync"

	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	"github.com/jonashiltl/sessions-backend/services/notification/internal/handler"
	gonats "github.com/nats-io/nats.go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	opts := []gonats.Option{gonats.Name("Notification Service")}
	nc, err := nats.Connect(opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	s := handler.NewServer()

	var wg sync.WaitGroup
	wg.Add(6)

	go func() {
		_, err := nc.QueueSubscribe("notification.email.verification", "notification.email.verification", s.EmailVerification)
		if err != nil {
			log.Fatal("Email Verification couldn't start", err)
		}
	}()
	go func() {
		_, err := nc.QueueSubscribe("notification.push.party.created", "notification.push.party.created", s.PartyCreated)
		if err != nil {
			log.Fatal("Party Created couldn't start", err)
		}
	}()
	go func() {
		_, err := nc.QueueSubscribe("notification.push.comment.created", "notification.push.comment.created", s.Commented)
		if err != nil {
			log.Fatal("Comment created couldn't start", err)
		}
	}()
	go func() {
		_, err := nc.QueueSubscribe("notification.push.comment.replied", "notification.push.comment.replied", s.Replied)
		if err != nil {
			log.Fatal("Replied couldn't start", err)
		}
	}()
	go func() {
		_, err := nc.QueueSubscribe("notification.push.friend.requested", "notification.push.friend.requested", s.FriendRequest)
		if err != nil {
			log.Fatal("Friend Request couldn't start", err)
		}
	}()
	go func() {
		_, err := nc.QueueSubscribe("notification.push.friend.accepted", "notification.push.friend.accepted", s.FriendAccepted)
		if err != nil {
			log.Fatal("Friend Accepted couldn't start", err)
		}
	}()

	// this will wait until the wg counter is at 0
	wg.Wait()
}
