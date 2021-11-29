package main

import (
	"flag"
	datastruct "local/MX-Steve/projects/grpc-demo/data-struct"
	pb "local/MX-Steve/projects/grpc-demo/data-struct/proto/server/djentry"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	config   = flag.String("conf", "", "The location of the config file.")
	endpoint = flag.String("endpoint", "127.0.0.1:50051", "The endpoint that the gRPC server should listen to")
)

func main() {
	flag.Parse()
	server, err := datastruct.NewStatisticsServer(*config)
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", *endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDjentryStatisticsServer(s, server)
	log.Infof("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Infof("failed to serve: %v", err)
	}
}
