package main

import (
	"WB_L0/internal/app"
	"WB_L0/internal/db"
	"WB_L0/internal/publisher"
	"WB_L0/internal/subscriber"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/stan.go"
	"io"
	"log"
	"os"
	"sync"
)

func main() {

	const connString = "postgres://me:123@localhost:5432/wbdb"
	const NATSClusterID = "wbcluster"
	const NATSClientID = "wbclient"
	cfg, _ := pgx.ParseConfig(connString)
	ctx, _ := context.WithCancel(context.Background())
	logger := log.Default()

	pgConn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		logger.Fatalf("unable to connect to database: %v\n", err)
	}
	stanConn, err := stan.Connect(NATSClusterID, NATSClientID)
	defer func() {
		stanConn.Close()
	}()

	mainCh := "orderch"
	goCh := make(chan any)
	wbApp := app.NewApp(ctx, pgConn, "orders")

	p := publisher.NewPublisher(stanConn, mainCh)
	s := subscriber.NewSubscriber(stanConn, wbApp)
	s.Subscribe(mainCh, goCh)

	f, err := os.Open("example.json")
	if err != nil {
		fmt.Println(err)
	}
	data, _ := io.ReadAll(f)
	var order db.OrderModel
	json.Unmarshal(data, &order)
	_ = p.Publish(p.Chan, data)

	//
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ch chan any) {
		defer wg.Done()
		for {
			<-ch
			gotData, err := wbApp.GetOrder(uuid.MustParse(order.OrderUID))
			if err != nil {
				panic(err)
			}
			fmt.Println(string(gotData))
		}
	}(goCh)
	wg.Wait()
	pgConn.Close(ctx)
}
