package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const CHUNK_BITS = 20

type FileMeta struct {
	Type   string
	Hash   string
	Name   string
	Size   int64
	Mode   string
	Chunks []string
}

func (fm *FileMeta) String() string {
	b, err := json.MarshalIndent(fm, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func BackupFile(bc *BackupClient, path string) (hash string, err error) {
	fmt.Println(path)
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return "", err
	}
	filemeta := FileMeta{
		Type: "file",
		Name: stat.Name(),
		Size: stat.Size(),
		Mode: stat.Mode().String(),
	}

	filesize := uint64(0)
	fullhash := sha256.New()
	chunk := []byte{}
	sum := MovingSum{}
	reader := bufio.NewReader(f)
	for {
		// Read file one byte at a time and split into chunks
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				// Perform cleanup
				fullhash.Write(chunk)
				filesize += uint64(len(chunk))
				hash, err := bc.PutChunk(chunk)
				if err != nil {
					return "", err
				}
				filemeta.Chunks = append(filemeta.Chunks, hash)
				break
			} else {
				return "", err
			}
		}
		sum.Add(b)
		chunk = append(chunk, b)
		if sum.OnSplit(18) && len(chunk) > 8194 {
			fullhash.Write(chunk)
			filesize += uint64(len(chunk))
			hash, err := bc.PutChunk(chunk)
			if err != nil {
				return "", err
			}
			filemeta.Chunks = append(filemeta.Chunks, hash)
			chunk = []byte{}
		}
	}
	filemeta.Hash = fmt.Sprintf("sha256:%x", fullhash.Sum(nil))
	fmt.Printf("%+V\n", filemeta)
	return bc.PutChunk([]byte(filemeta.String()))
}
