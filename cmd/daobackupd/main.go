package main

import (
	"log"
	"net"
	"os"

	"github.com/floort/daobackup"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	server := daobackup.BasicFilesystemServer{Root: os.Args[1]}
	grpcServer := grpc.NewServer(opts...)
	daobackup.RegisterDAOBackupServer(grpcServer, &server)
	grpcServer.Serve(lis)
}
