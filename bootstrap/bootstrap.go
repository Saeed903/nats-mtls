package bootstrap

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

// CreateJStream create stream with subjects in nats-server
func CreateJStream(streamName, streamSubject string ,nc *nats.Conn) error {
	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("jetstream err: %v",err)
	}

	// check if the CAMERAS stream aleady exists; if not, create it.
	stream, err := js.StreamInfo(streamName)
	if err != nil {
		log.Printf("stream not created: %v",err)
	}
	if stream == nil {
		log.Printf("creating stream %s and subjects %s", streamName, streamSubject)
		_, err := js.AddStream(&nats.StreamConfig{
			Name: streamName,
			Subjects: []string{streamSubject},
		})
		if err != nil {
			return err
		}
	} else {
		fmt.Printf("stream %s and subjects %s existed.", streamName, streamSubject)
	}
	return nil
}