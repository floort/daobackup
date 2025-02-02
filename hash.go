package daobackup

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type ChunkHash [sha256.Size]byte

func (h ChunkHash) String() string {
	return fmt.Sprintf("%x", [32]byte(h))
}

func (h ChunkHash) Bytes() []byte {
	return h[:]
}

func (h ChunkHash) Verify(blob []byte) bool {
	newhash := sha256.Sum256(blob)
	if len(h) != len(newhash) {
		return false
	}
	return ChunkHash(newhash) == h
}

func ChunkHashFromBytes(h []byte) ChunkHash {
	if len(h) != len(ChunkHash{}) {
		return ChunkHash{}
	}
	return ChunkHash(h)
}

func ParseChunkHash(hashstring string) (h ChunkHash, err error) {
	decoded, err := hex.DecodeString(hashstring)
	if err != nil {
		return ChunkHash{}, err
	}
	if len(decoded) != len(ChunkHash{}) {
		return ChunkHash{}, fmt.Errorf("could not decode '%s' as a %d byte hash", hashstring, len(ChunkHash{}))
	}
	return ChunkHash(decoded), nil
}

func IsValidHash(hashstring string) bool {
	_, err := ParseChunkHash(hashstring)
	return err == nil
}

func HashChunk(chunk []byte) (hash ChunkHash) {
	return sha256.Sum256(chunk)
}
