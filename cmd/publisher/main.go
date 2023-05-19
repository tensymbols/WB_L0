package main

import (
	"WB_L0/internal/publisher"
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"log"
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

	f, err := os.Open("../example.json") // reading example order
	if err != nil {
		fmt.Println(err)
	}
	data, _ := io.ReadAll(f)
	log.Println("publishing order...")
	_ = p.Publish(p.Chan, data) // publishing example order
	fmt.Println("press any key to publish garbage")
	fmt.Scanln()
	log.Println()
	_ = p.Publish(p.Chan, []byte("garbage")) // publishing invalid data
}
