package stream

import (
	"reflect"
	"strings"
	"unicode"

	"github.com/nats-io/nats.go"
)

type stream struct {
	nats *nats.EncodedConn
}

type Stream interface {
	PublishEvent(event interface{}) error
}

func NewStream(nats *nats.EncodedConn) Stream {
	return &stream{nats: nats}
}

func (s stream) PublishEvent(event interface{}) error {
	sub := eventToSubject(event)
	return s.nats.Publish(sub, event)
}

func (s stream) SubscribeByEvent() {

}

func eventToSubject(event interface{}) string {
	t := reflect.TypeOf(event)

	s := strings.Split(t.String(), ".")

	// if type is events.ImportantType, remove events prefix from string
	if len(s) == 2 && s[0] == "events" {
		return camelcaseStringToDotString(s[1])
	}

	return camelcaseStringToDotString(t.String())
}

func camelcaseStringToDotString(camelcase string) string {
	var b strings.Builder

	for i, c := range camelcase {
		if unicode.IsUpper(c) {
			if i != 0 {
				b.WriteString(".")
			}
			b.WriteRune(unicode.ToLower(c))
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}
