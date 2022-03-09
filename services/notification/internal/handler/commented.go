package handler

import (
	"log"
	"os"

	gonats "github.com/nats-io/nats.go"
)

func (*server) Commented(m *gonats.Msg) {
	log.Printf("Received on [%s] Queue[%s] Pid[%d]: '%s'", m.Subject, m.Sub.Queue, os.Getpid(), string(m.Data))
}
