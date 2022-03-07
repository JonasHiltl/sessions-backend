package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/jonashiltl/sessions-backend/packages/nats"
	gonats "github.com/nats-io/nats.go"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	nc, err := nats.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	go nc.QueueSubscribe("notification.email.verification", "email_verification", func(m *gonats.Msg) {
		log.Println("Email Verification Notification received")
		log.Println("Data", m.Data)
	})
	go nc.QueueSubscribe("notification.push.party.created", "party_created", func(m *gonats.Msg) {
		log.Println("Party Creation Notification received")
		log.Println("Data", m.Data)
	})
	go nc.QueueSubscribe("notification.push.comment.created", "party_comment", func(m *gonats.Msg) {
		log.Println("Comment On Party Notification received")
		log.Println("Data", m.Data)
	})
	go nc.QueueSubscribe("notification.push.comment.replied", "comment_reply", func(m *gonats.Msg) {
		log.Println("Reply on Comment Notification received")
		log.Println("Data", m.Data)
	})
	go nc.QueueSubscribe("notification.push.friend.requested", "friend_request", func(m *gonats.Msg) {
		log.Println("Friend Request Notification received")
		log.Println("Data", m.Data)
	})
	go nc.QueueSubscribe("notification.push.friend.accepted", "friend_accepted", func(m *gonats.Msg) {
		log.Println("Friend Request Accepted Notification received")
		log.Println("Data", m.Data)
	})
}
