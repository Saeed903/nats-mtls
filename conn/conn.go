package conn

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/saeed903/nats-mtls/constants"
)


func Connect() (*nats.Conn, error) {
	nc, err := nats.Connect(
		fmt.Sprintf("tls://%s:%d", constants.Host, constants.Port),
		nats.RootCAs(constants.CaFile),
		nats.ClientCert(constants.ClientCertFile, constants.ClientKeyFile),
	)
	if err != nil {
		log.Fatalf("client connect: %v", err)
		return nil, err
	}
	return nc, nil
}