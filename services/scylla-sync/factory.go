package main

import (
	"context"
	"time"

	"github.com/gocql/gocql"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/scylla-sync/consumer"
	scyllacdc "github.com/scylladb/scylla-cdc-go"
)

type Factory interface {
	CreateChangeConsumer(ctx context.Context, input scyllacdc.CreateChangeConsumerInput) (scyllacdc.ChangeConsumer, error)
}

type factory struct {
	logger  scyllacdc.Logger
	session *gocql.Session

	stream stream.Stream

	progressReporterInterval time.Duration
	rowsRead                 *int64
}

func (f *factory) CreateChangeConsumer(ctx context.Context, input scyllacdc.CreateChangeConsumerInput) (scyllacdc.ChangeConsumer, error) {
	reporter := scyllacdc.NewPeriodicProgressReporter(f.logger, f.progressReporterInterval, input.ProgressReporter)
	reporter.Start(ctx)
	return consumer.NewConsumer(f.stream, reporter), nil
}
