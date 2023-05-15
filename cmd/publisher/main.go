package main

import (
	"WB_L0/internal/publisher"
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"os"
)

const (
	stanClusterID = "wbcluster"
	stanClientID  = "publisher"
	channelName   = "orderch"
)

func main() {
	sc, err := stan.Connect(stanClusterID, stanClientID)
	if err != nil {
		panic(fmt.Sprintf("can't connect to NATS streaming server: %v ", err))
	}
	defer sc.Close()
	p := publisher.NewPublisher(sc, channelName)

	f, err := os.Open("../example.json")
	if err != nil {
		fmt.Println(err)
	}
	data, _ := io.ReadAll(f)

	_ = p.Publish(p.Chan, data)
	_ = p.Publish(p.Chan, []byte("garbage"))
}
