package main

import (
	"log"
	"strconv"

	"github.com/saeed903/nats-mtls/bootstrap"
	"github.com/saeed903/nats-mtls/conn"
	"github.com/saeed903/nats-mtls/constants"
)

func main() {
	nc, err := conn.Connect()
	if err != nil {
		log.Fatalf("client connect: %v", err)
	}

	err = bootstrap.CreateJStream(constants.CameraStream, constants.CameraStreamSubjects, nc)
	if err != nil {
		log.Fatalf("CreateJStream err: %v", err)
	}

	jsc, err := nc.JetStream()
	for i:=1; i<1000000; i++{
	jsc.Publish(constants.CameraStreamSubjects, []byte("hello sadrsys" + strconv.Itoa(i)))
	log.Println("published message on subject: " + constants.CameraStreamSubjects)
	}

	// nc.Publish(constants.CameraStreamSubjects, []byte("world"))
	// nc.Publish(constants.CameraStreamSubjects, []byte("Saeed is start mtls running..."))
	// nc.Publish(constants.CameraStreamSubjects, []byte("Best Camera Security running..."))
	// nc.Publish(constants.CameraStreamSubjects, []byte("Sadra myson ..."))
}