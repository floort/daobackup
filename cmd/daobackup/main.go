package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/floort/daobackup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(":8000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := daobackup.NewDAOBackupClient(conn)

	h := daobackup.ChunkHash{}
	c := daobackup.ErrorBytes{}
	for h, c = range daobackup.ChunkDir(os.Args[1]) {
		if c.Error != nil {
			panic(c.Error)
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		status, err := client.CheckBlob(ctx, &daobackup.Hash{Hash: h.Bytes()})
		if err != nil {
			panic(err)
		}
		if status.Succes {
			fmt.Printf("Hash %s already available.", h.String())
		} else {
			status, err = client.PutBlob(ctx, &daobackup.HashedBlob{Hash: h.Bytes(), Blob: c.Bytes})
			if err != nil {
				panic(err)
			}
			if !status.Succes {
				fmt.Print(status.Message)
			}
		}
	}
}
