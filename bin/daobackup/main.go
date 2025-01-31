package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/floort/daobackup"
)

func main() {
	h := daobackup.ChunkHash{}
	c := daobackup.ErrorBytes{}
	for h, c = range daobackup.ChunkDir(os.Args[1]) {
		if c.Error != nil {
			panic(c.Error)
		}
		client := http.Client{}
		r, err := http.NewRequest(http.MethodPut, "http://localhost:8080/chunk/"+h.String(), bytes.NewReader(c.Bytes))
		if err != nil {
			panic(err)
		}
		resp, err := client.Do(r)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != http.StatusCreated {
			panic(resp.Status)
		}
	}
}
