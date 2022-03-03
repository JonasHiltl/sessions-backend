package consumer

import (
	"errors"
	"os"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/encoders/protobuf"
)

func Connect() (*nats.EncodedConn, error) {
	cluster, exists := os.LookupEnv("NATS_CLUSTER")
	if !exists {
		return nil, errors.New("nats connection url not defined")
	}

	nc, err := nats.Connect(cluster)
	if err != nil {
		return nil, err
	}
	c, err := nats.NewEncodedConn(nc, protobuf.PROTOBUF_ENCODER)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func Start(nc *nats.EncodedConn) {

}
