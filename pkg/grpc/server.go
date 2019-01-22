package grpc

import (
	"fmt"
	"github.com/batazor/go-smpp/pb"
	"github.com/batazor/go-smpp/pkg/utils"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var (
	log       = logrus.New()
	GRPC_PORT = utils.Getenv("GRPC_PORT", "50051")
)

func init() {
	// Logging =================================================================
	// Setup the logger backend using Sirupsen/logrus and configure
	// it to use a custom JSONFormatter. See the logrus docs for how to
	// configure the backend at github.com/Sirupsen/logrus
	log.Formatter = new(logrus.JSONFormatter)
}

func Listen() {
	port := fmt.Sprintf(":%s", GRPC_PORT)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error open port for gRPC: ", err)
	}

	log.Info("run gRPC on port " + port)

	// Enable Prometheus histograms
	grpc_prometheus.EnableHandlingTimeHistogram()

	// Initialize gRPC server's interceptor
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		grpc.UnaryInterceptor(grpc_prometheus.UnaryServerInterceptor),
	)

	// Register gRPC service implementations
	command.RegisterCommandServer(grpcServer, &server{})

	// After all your registrations, make sure all of the Prometheus metrics are initialized.
	grpc_prometheus.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
