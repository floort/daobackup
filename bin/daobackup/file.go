package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aclements/go-rabin/rabin"
    "github.com/kr/pretty"
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

	chunker := rabin.NewChunker(rabin.NewTable(rabin.Poly64, 512), f, 1024, 8*1024*1024, 32*1024*1024)
	fullhash := sha256.New()
	for chunksize, err := chunker.Next(); err == nil; chunksize, err = chunker.Next() {
        if chunksize == 0 {
            break
        }
		// Rewind file to start of the chunk
		f.Seek(-int64(chunksize), os.SEEK_CUR)
		// Read chunk
		buf := make([]byte, chunksize)
		n, err := f.Read(buf)
		if err != nil {
			return "", err
		}
		if n != chunksize {
			panic("buffer error")
		}
		fmt.Println(fullhash.Write(buf))
		hash, err := bc.PutChunk(buf)
		if err != nil {
			return "", err
		}
		filemeta.Chunks = append(filemeta.Chunks, hash)
	}
	filemeta.Hash = fmt.Sprintf("sha256:%x", fullhash.Sum(nil))
	fmt.Printf("%# v\n", pretty.Formatter(filemeta))
	return bc.PutChunk([]byte(filemeta.String()))

}
