package main

import (
	"WB_L0/internal/ports"
	"WB_L0/internal/service"
	"github.com/jackc/pgx/v5"
	"github.com/nats-io/stan.go"
	"golang.org/x/net/context"
	"log"
)

func startServer(s service.Service) *ports.Server {
	server := ports.NewServer(serverPort, s)
	log.Println("starting server")
	go server.Listen()
	log.Println("server has started")
	return server
}

func stanConnect() stan.Conn {
	log.Printf("connecting to NATS Streaming server: %s as %s", stanClusterID, stanClientID)
	stanConn, err := stan.Connect(stanClusterID, stanClientID)
	if err != nil {
		log.Fatalf("unable to connect to NATS Streaming server: %v\n", err)

	}
	return stanConn
}

func dbConnect(ctx context.Context) *pgx.Conn {
	cfg, _ := pgx.ParseConfig(connString)
	log.Println("connecting to database")
	c, err := pgx.ConnectConfig(ctx, cfg)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}

	return c
}

func startService(ctx context.Context, conn *pgx.Conn) service.Service {
	log.Println("starting service")
	return service.NewService(ctx, conn, "orders")
}
