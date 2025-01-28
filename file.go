package daobackup

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path"
)

type ErrorBytes struct {
	Error error
	Bytes []byte
}

func ChunkFile(path string) func(func(ChunkHash, ErrorBytes) bool) {
	// Setup
	filehandle, err := os.Open(path)
	if err != nil {
		return func(yield func(ChunkHash, ErrorBytes) bool) {
			yield(ChunkHash{}, ErrorBytes{err, nil})
		}
	}
	// Fill in the metadata
	filemeta := FileMeta{}
	// Fill in parts of the metadata
	stat, err := filehandle.Stat()
	if err != nil {
		return func(yield func(ChunkHash, ErrorBytes) bool) {
			yield(ChunkHash{}, ErrorBytes{err, nil})
			filehandle.Close()
		}
	}
	if stat.IsDir() {
		return func(yield func(ChunkHash, ErrorBytes) bool) {
			yield(ChunkHash{}, ErrorBytes{fmt.Errorf("ChunkFile can't chunk a directory"), nil})
		}
	}
	filemeta.Type = Normal
	filemeta.Name = stat.Name()
	filemeta.Size = stat.Size()
	filemeta.ModTime = stat.ModTime()
	filemeta.Mode = uint32(stat.Mode())

	hasher := sha256.New()
	chunkbuffer := make([]byte, MaxChunkSize)
	chunker := NewChunker(filehandle)
	return func(yield func(ChunkHash, ErrorBytes) bool) {
		for {
			chunk, err := chunker.Next(chunkbuffer)
			if err == io.EOF { // No more data chunks, return FileMeta chunk
				filemeta.Hash = ChunkHash(hasher.Sum([]byte{}))
				bytesbuf := make([]byte, 0)
				bytesbuf, err := filemeta.MarshalMsg(bytesbuf)
				if err != nil {
					yield(ChunkHash{}, ErrorBytes{err, nil})
					filehandle.Close()
					return
				}
				yield(Hash(bytesbuf), ErrorBytes{nil, bytesbuf})
				filehandle.Close()
				return
			}
			if err != nil {
				yield(ChunkHash{}, ErrorBytes{err, nil})
				filehandle.Close()
				return
			}
			chunkhash := Hash(chunk.Data)
			filemeta.Chunks = append(filemeta.Chunks, ChunkMeta{Offset: uint64(chunk.Start), Size: uint64(chunk.Length), Hash: chunkhash})
			// update total file hash
			hasher.Write(chunk.Data)
			if !yield(chunkhash, ErrorBytes{nil, chunk.Data}) {
				filehandle.Close()
				return
			}
		}
	}
}

func ChunkDir(dirpath string) func(func(ChunkHash, ErrorBytes) bool) {
	// Setup
	filehandle, err := os.Open(dirpath)
	if err != nil {
		return func(yield func(ChunkHash, ErrorBytes) bool) {
			yield(ChunkHash{}, ErrorBytes{err, nil})
		}
	}
	defer filehandle.Close()
	// Fill in the metadata
	filemeta := FileMeta{}
	// Fill in parts of the metadata
	stat, err := filehandle.Stat()
	if err != nil {
		return func(yield func(ChunkHash, ErrorBytes) bool) {
			yield(ChunkHash{}, ErrorBytes{err, nil})
		}
	}
	if !stat.IsDir() {
		return func(yield func(ChunkHash, ErrorBytes) bool) {
			yield(ChunkHash{}, ErrorBytes{fmt.Errorf("ChunkDir can't chunk a file"), nil})
		}
	}
	filemeta.Type = Directory
	filemeta.Name = stat.Name()
	filemeta.Size = stat.Size()
	filemeta.ModTime = stat.ModTime()
	filemeta.Mode = uint32(stat.Mode())
	filemeta.Entries = map[string]ChunkHash{}
	return func(yield func(ChunkHash, ErrorBytes) bool) {
		files, err := os.ReadDir(dirpath)
		if err != nil {
			yield(ChunkHash{}, ErrorBytes{err, nil})
			return
		}
		for _, file := range files {
			fullpath := path.Join(dirpath, file.Name())
			hash := ChunkHash{}
			blob := ErrorBytes{}
			if file.IsDir() {
				for hash, blob = range ChunkDir(fullpath) {
					if !yield(hash, blob) {
						break
					}
				}
				filemeta.Entries[file.Name()] = hash
			} else {
				for hash, blob := range ChunkFile(fullpath) {
					if !yield(hash, blob) {
						break
					}
				}
				filemeta.Entries[file.Name()] = hash
			}

		}
		bytesbuf := make([]byte, 0)
		bytesbuf, err = filemeta.MarshalMsg(bytesbuf)
		if err != nil {
			yield(ChunkHash{}, ErrorBytes{err, nil})
			return
		}
		yield(Hash(bytesbuf), ErrorBytes{nil, bytesbuf})
	}
}
