package main

import (
	"WB_L0/internal/subscriber"
	"context"
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

	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	pgConn := dbConnect(ctx)
	defer pgConn.Close(ctx)
	stanConn := stanConnect()
	defer stanConn.Close()
	wbService := startService(ctx, pgConn)

	startServer(wbService)

	sub := subscriber.NewSubscriber(stanConn, wbService, "testsub")
	sub.Subscribe(channelName)
	<-ctx.Done()
}
