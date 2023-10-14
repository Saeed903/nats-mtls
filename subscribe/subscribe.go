package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/saeed903/nats-mtls/conn"
	"github.com/saeed903/nats-mtls/constants"
)





func main() {
	nc, err := conn.Connect()
	if err != nil {
		log.Fatalf("nats connect: %v", err)
	}

	// err = bootstrap.CreateJStream(constants.CameraStream, constants.CameraStreamSubjects, nc)
	// if err != nil {
	// 	log.Fatalf("CreateJStream err: %v", err)
	// }

	jsc, err := nc.JetStream()

	// _, err = jsc.Subscribe(constants.CameraStreamSubjects, func(msg *nats.Msg) {
	// 	err := msg.Ack()

	// 	if err != nil {
	// 		log.Println("Unable to Ack", err)
	// 		return
	// 	}
	// 	log.Printf("Consumer => Subject: %s   - data: %s", constants.CameraStreamSubjects, string(msg.Data))
	// })

	// if err != nil {
	// 	log.Println("Subscribe failed")
	// 	return
	// }

	
	sub, _ := jsc.PullSubscribe(constants.CameraStreamSubjects, "", nats.PullMaxWaiting(128))
	for {
		msgs, _ := sub.Fetch(constants.Batch)
		for _, msg := range msgs {
			msg.Ack()
			log.Println(string(msg.Data))
		}
	}

	// sub, _ := jsc.SubscribeSync(constants.CameraStreamSubjects)
	// for {
	// 	msgs, _ := sub.Fetch(constants.Batch)
	// 	for _, msg := range msgs {
	// 		msg.Ack()
	// 		log.Println(string(msg.Data))
	// 	}
	// }
	// sub, err := nc.SubscribeSync("hello")
	// if err != nil {
	// 	log.Fatalf("subscribeSync error: %v", err)
	// }
	// defer sub.Drain()

	//time.Sleep(20 * time.Second)

	// msg, err := sub.NextMsg(5*time.Second)
	// if err != nil {
	// 	log.Fatalf("Next msg: %v", err)
	// }

	// fmt.Println(string(msg.Data))

	//select{}
}


