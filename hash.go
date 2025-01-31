package daobackup

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type ChunkHash [32]byte

func (h ChunkHash) String() string {
	return fmt.Sprintf("%x", [32]byte(h))
}

func ParseChunkHash(hashstring string) (h ChunkHash, err error) {
	decoded, err := hex.DecodeString(hashstring)
	if err != nil {
		return ChunkHash{}, err
	}
	if len(decoded) != len(ChunkHash{}) {
		return ChunkHash{}, fmt.Errorf("Could not decode '%s' as a %d byte hash", hashstring, len(ChunkHash{}))
	}
	return ChunkHash(decoded), nil
}

func IsValidHash(hashstring string) bool {
	_, err := ParseChunkHash(hashstring)
	return err == nil
}

func Hash(chunk []byte) (hash ChunkHash) {
	return sha256.Sum256(chunk)
}
