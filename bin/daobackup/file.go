package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/floort/daobackup/rollsum"
)

const CHUNK_BITS = 22

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
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
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
	chunks, err := chunkfile(f)
	if err != nil {
		return "", err
	}
	for c := range chunks {
		fullhash.Write(chunks[c])
		filesize += uint64(len(chunks[c]))
		hash, err := bc.PutChunk(chunks[c])
		if err != nil {
			return "", err
		}
		filemeta.Chunks = append(filemeta.Chunks, hash)
	}
	filemeta.Hash = fmt.Sprintf("sha256:%x", fullhash.Sum(nil))
	return bc.PutChunk([]byte(filemeta.String()))
}

func chunkfile(file *os.File) (chunks [][]byte, err error) {
	chunk := []byte{}
	sum := rollsum.New()
	reader := bufio.NewReader(file)
	for {
		// Read file one byte at a time and split into chunks
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				// Perform cleanup
				chunks = append(chunks, chunk)
				err = nil // End of file is not an error
			}
			return chunks, err
		}
		sum.Roll(b)
		chunk = append(chunk, b)
		if sum.OnSplitWithBits(CHUNK_BITS) {
			chunks = append(chunks, chunk)
			chunk = []byte{}
		}

	}
}
