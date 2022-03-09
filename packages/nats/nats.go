package nats

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
)

func Connect(opts []nats.Option) (*nats.EncodedConn, error) {
	cluster, exists := os.LookupEnv("NATS_CLUSTER")
	if !exists {
		return nil, errors.New("nats connection url not defined")
	}

	opts = setupConnOptions(opts)

	nc, err := nats.Connect(cluster, opts...)
	if err != nil {
		return nil, err
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to Nats Server at ", c.Conn.ConnectedUrl())
	return c, nil
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	totalWait := 10 * time.Minute
	reconnectDelay := time.Second

	opts = append(opts, nats.ReconnectWait(reconnectDelay))
	opts = append(opts, nats.MaxReconnects(int(totalWait/reconnectDelay)))
	opts = append(opts, nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s, will attempt reconnects for %.0fm", err, totalWait.Minutes())
	}))
	opts = append(opts, nats.ReconnectHandler(func(nc *nats.Conn) {
		log.Printf("Reconnected [%s]", nc.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(nc *nats.Conn) {
		log.Fatalf("Exiting: %v", nc.LastError())
	}))
	return opts
}
