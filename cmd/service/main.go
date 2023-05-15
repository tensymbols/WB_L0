package main

import (
	"WB_L0/internal/ports"
	"WB_L0/internal/service"
	"WB_L0/internal/subscriber"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/stan.go"
	"log"
	"os/signal"
	"syscall"
)

const (
	connString    = "postgres://me:123@localhost:5432/wbdb"
	stanClusterID = "wbcluster"
	stanClientID  = "wbclient"
	channelName   = "orderch"

	serverPort = ":8080"
)

func main() {

	cfg, _ := pgx.ParseConfig(connString)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	logger := log.Default()

	log.Println("connecting to database")
	pgConn, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		logger.Fatalf("unable to connect to database: %v\n", err)
		return
	}
	defer pgConn.Close(ctx)
	log.Printf("connecting to NATS Streaming server: %s as %s", stanClusterID, stanClientID)
	stanConn, err := stan.Connect(stanClusterID, stanClientID)
	if err != nil {
		logger.Fatalf("unable to connect to NATS Streaming server: %v\n", err)
		return
	}
	defer stanConn.Close()

	log.Println("starting service")
	wbService := service.NewService(ctx, pgConn, "orders")

	server := ports.NewServer(serverPort, wbService)
	log.Println("starting server")
	go server.Listen()
	log.Println("server has started")
	sub := subscriber.NewSubscriber(stanConn, wbService, "testsub")
	sub.Subscribe(channelName)
	<-ctx.Done()
}
