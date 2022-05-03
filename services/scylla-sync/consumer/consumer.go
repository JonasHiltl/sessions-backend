package consumer

import (
	"context"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/jonashiltl/sessions-backend/packages/events"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	scyllacdc "github.com/scylladb/scylla-cdc-go"
)

type Consumer interface {
	Consume(ctx context.Context, change scyllacdc.Change) error
	End() error
	PartyCreated(cr *scyllacdc.ChangeRow) error
}

type consumer struct {
	stream   stream.Stream
	reporter *scyllacdc.PeriodicProgressReporter
}

func NewConsumer(stream stream.Stream, reporter *scyllacdc.PeriodicProgressReporter) Consumer {
	return &consumer{
		stream:   stream,
		reporter: reporter,
	}
}

func (c *consumer) Consume(ctx context.Context, change scyllacdc.Change) error {
	//if tableName == "sessions.parties" {
	for _, cr := range change.Delta {
		err := c.PartyCreated(cr)
		if err != nil {
			return err
		}
	}
	//}

	c.reporter.Update(change.Time)
	return nil
}

func (c *consumer) End() error {
	_ = c.reporter.SaveAndStop(context.Background())
	return nil
}

func (c consumer) PartyCreated(cr *scyllacdc.ChangeRow) error {
	idRaw, _ := cr.GetValue("id")
	id, ok := idRaw.(*gocql.UUID)
	if !ok {
		log.Println("Failed to cast id")
	}

	uIdRaw, _ := cr.GetValue("user_id")
	uId, ok := uIdRaw.(*string)
	if !ok {
		log.Println("Failed to cast user id")
	}

	titleRaw, _ := cr.GetValue("title")
	title, ok := titleRaw.(*string)
	if !ok {
		log.Println("Failed to title")
	}

	isPublicRaw, _ := cr.GetValue("is_public")
	isPublic, ok := isPublicRaw.(*bool)
	if !ok {
		log.Println("Failed to cast isPublic")
	}

	geohashRaw, _ := cr.GetValue("geohash")
	geohash, ok := geohashRaw.(*string)
	if !ok {
		log.Println("Failed to geohash")
	}

	latRaw, _ := cr.GetValue("lat")
	lat, ok := latRaw.(*float32)
	if !ok {
		log.Println("Failed to cast lat")
	}

	longRaw, _ := cr.GetValue("long")
	long, ok := longRaw.(*float32)
	if !ok {
		log.Println("Failed to cast long")
	}

	streetAddressRaw, _ := cr.GetValue("street_address")
	streetAddress, ok := streetAddressRaw.(*string)
	if !ok {
		log.Println("Failed to cast street address")
	}

	postalCodeRaw, _ := cr.GetValue("postal_code")
	postalCode, ok := postalCodeRaw.(*string)
	if !ok {
		log.Println("Failed to cast postal code")
	}

	stateRaw, _ := cr.GetValue("state")
	state, ok := stateRaw.(*string)
	if !ok {
		log.Println("Failed to cast state")
	}

	countryRaw, _ := cr.GetValue("country")
	country, ok := countryRaw.(*string)
	if !ok {
		log.Println("Failed to cast country")
	}

	startDateRaw, _ := cr.GetValue("start_date")
	startDate, ok := startDateRaw.(*time.Time)
	if !ok {
		log.Println("Failed to cast startDate")
	}

	e := &events.PartyCreated{
		Id:            id.String(),
		UserId:        nullableStringToStr(uId),
		Title:         nullableStringToStr(title),
		IsPublic:      nullableBoolToBool(isPublic),
		Geohash:       nullableStringToStr(geohash),
		Lat:           nullableFloatToFloat(lat),
		Long:          nullableFloatToFloat(long),
		StreetAddress: nullableStringToStr(streetAddress),
		PostalCode:    nullableStringToStr(postalCode),
		State:         nullableStringToStr(state),
		Country:       nullableStringToStr(country),
		StartDate:     nullableTimeToStr(startDate),
		CreatedAt:     id.Time().UTC().Format(time.RFC3339),
	}

	log.Println(e)

	err := c.stream.PublishEvent(e)
	if err != nil {
		log.Println(err)
	}

	return nil
}
