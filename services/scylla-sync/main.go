package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gocql/gocql"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/scylla-sync/config"
	"github.com/nats-io/nats.go"
	scyllacdc "github.com/scylladb/scylla-cdc-go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	cluster := gocql.NewCluster(c.SCYLLA_HOSTS)
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.DCAwareRoundRobinPolicy("local-dc"))
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	opts := []nats.Option{nats.Name("Scylla Sync Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	stream := stream.New(nc)
	logger := log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	rowsRead := new(int64)

	var progressManager scyllacdc.ProgressManager
	progressManager, err = scyllacdc.NewTableBackedProgressManager(session, "sessions.scylla_sync_progress", "cdc-replicator")
	if err != nil {
		log.Fatalln(err)
	}

	factory := &factory{
		session: session,
		stream:  stream,

		logger: logger,

		progressReporterInterval: time.Second * 30,
		rowsRead:                 rowsRead,
	}

	adv := scyllacdc.AdvancedReaderConfig{
		ConfidenceWindowSize:   10 * time.Second,
		ChangeAgeLimit:         10 * time.Minute,
		QueryTimeWindowSize:    5 * 60 * time.Second,
		PostEmptyQueryDelay:    5 * time.Second,
		PostNonEmptyQueryDelay: 5 * time.Second,
		PostFailedQueryDelay:   1 * time.Second,
	}

	cfg := &scyllacdc.ReaderConfig{
		Session:               session,
		ChangeConsumerFactory: factory,
		ProgressManager:       progressManager,
		TableNames:            c.SCYLLA_TABLES,
		Logger:                logger,
		Advanced:              adv,
	}

	reader, err := scyllacdc.NewReader(context.Background(), cfg)
	if err != nil {
		log.Fatalln(err)
	}

	signalC := make(chan os.Signal)
	go func() {
		<-signalC
		reader.Stop()

		<-signalC
		os.Exit(1)
	}()
	signal.Notify(signalC, os.Interrupt)

	if err := reader.Run(context.Background()); err != nil {
		log.Fatal(err)
	}

}
