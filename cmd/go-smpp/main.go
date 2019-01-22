package main

import (
	"github.com/batazor/go-smpp/pkg/grpc"
	"github.com/batazor/go-smpp/pkg/mongo"
	"github.com/batazor/go-smpp/pkg/smpp"
	"github.com/batazor/go-smpp/pkg/utils"
	"github.com/sirupsen/logrus"
)

var (
	// Logging
	log = logrus.New()

	// ENV
	SMPP_ENABLE  = utils.Getenv("SMPP_ENABLE", "true")
	GRPC_ENABLE  = utils.Getenv("GRPC_ENABLE", "true")
	MONGO_ENABLE = utils.Getenv("MONGO_ENABLE", "true")
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func main() {
	// Run SMPP
	if SMPP_ENABLE == "true" {
		go smpp.Listen()
	}

	// Run gRPC
	if GRPC_ENABLE == "true" {
		go grpc.Listen()
	}

	// Run mongo
	if MONGO_ENABLE == "true" {
		go mongo.Listen()
	}

	select {}
}
