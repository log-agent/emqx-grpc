package main

import (
	"log"

	"github.com/log-agent/emqx-grpc/grpc"
)

var (
	Buildtime string
	Version   string
)

func main() {
	if err := grpc.GrpcRun(); err != nil {
		log.Println("grpc.GrpcRun", err.Error())
	}
}
