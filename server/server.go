package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/saeed903/nats-mtls/constants"
)

// var (
// 	host     string
// 	port     int
// 	certFile string
// 	keyFile  string
// 	caFile   string
// )


func main() {
	// flag.StringVar(&host, "host", "localhost", "Client connection host/ip.")
	// flag.IntVar(&port, "port", 4222, "Client connection port.")
	// flag.StringVar(&certFile, "tls.cert.server", "cert.pem", "TLS cert file.")
	// flag.StringVar(&keyFile, "tls.key.server", "key.pem", "TLS key file.")
	// flag.StringVar(&caFile, "tls.ca", "ca.pem", "TLS CA file.")

	// flag.Parse()

	serverTlsConfig, err := server.GenTLSConfig(&server.TLSConfigOpts{
		CertFile: constants.CertFile,
		KeyFile: constants.KeyFile,
		CaFile: constants.CaFile,
		Verify: true,
		Timeout: 2,
	})
	if err != nil {
		log.Fatalf("tls config: %v", err)
	}

	// Setup the embeded server options.
	opts := server.Options{
		TLSConfig: serverTlsConfig,
		Host: constants.Host,
		Port: constants.Port,
		JetStream: true,
		StoreDir: "/data/jetstream-server",
		JetStreamMaxMemory: 1073741824,
		JetStreamMaxStore: 1073741824,
	}

	// Initialize a new server with the options.
	ns, err := server.NewServer(&opts)
	if err != nil {
		log.Fatalf("server init: %v", err)
	}

	go ns.Start()
	defer ns.Shutdown()

	fmt.Print("Nats server is starting...")

	// var wg sync.WaitGroup

	// nc, err := conn.Connect()

	// err = bootstrap.CreateJStream(constants.CameraStream, constants.CameraStreamSubjects, nc)
	// if err != nil {
	// 	log.Fatalf("CreateJStream err: %v", err)
	// }

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()

	// 	jsc, _ := nc.JetStream()
	// 	jsc.Publish(constants.CameraStreamSubjects, []byte("hello sadrsys"))
	// 	log.Println("published message on subject: " + constants.CameraStreamSubjects)
	// }()

	// wg.Add(1)
	// go func(){
	// 	defer wg.Done()
	// 	jsc, _ := nc.JetStream()
	// 	_, err := jsc.Subscribe(constants.CameraStreamSubjects, func(msg *nats.Msg) {
	// 		err := msg.Ack()

	// 		if err != nil {
	// 			log.Println("Unable to Ack", err)
	// 			return
	// 		}
	// 		log.Printf("Consumer => Subject: %s   - data: %s", constants.CameraStreamSubjects, string(msg.Data))
	// 	})

	// 	if err != nil {
	// 		log.Println("Subscribe failed")
	// 		return
	// 	}
	// // 	sub, _ := jsc.SubscribeSync(constants.CameraStreamSubjects)
	// // 	for {
	// // 		msgs, _ := sub.Fetch(constants.Batch)
	// // 		for _, msg := range msgs {
	// // 			msg.Ack()
	// // 			log.Println(string(msg.Data))
	// // 		}
	// // }
	// }()
	// wg.Wait()

	select {}

}