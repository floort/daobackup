package main

import (
	"fmt"
	"os"
)

func main() {
	bc := BackupClient{}
	hash, err := BackupDir(&bc, os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", hash, os.Args[1])
}
