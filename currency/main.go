package main

import (
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	protos "github.com/moaabb/microservices/currency/protos/currency"
	"github.com/moaabb/microservices/currency/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":9092"

func main() {
	log := hclog.Default()

	gs := grpc.NewServer()
	cs := server.NewCurrency(log)
	protos.RegisterCurrencyServer(gs, cs)

	reflection.Register(gs)
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	log.Info("Server listening on port", port)
	err = gs.Serve(l)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
