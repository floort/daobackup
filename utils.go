package daobackup

import (
	"io"

	"github.com/restic/chunker"
)

const (
	MinChunkSize = 512 * 1024      // 512 KB
	MaxChunkSize = 8 * 1024 * 1024 // 8 MB
)

func NewChunker(rd io.Reader) *chunker.Chunker {
	return chunker.NewWithBoundaries(rd, chunker.Pol(0x3DA3358B4DC173), MinChunkSize, MaxChunkSize)

}
