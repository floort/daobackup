package main

import (
	"crypto/sha256"
	"fmt"
)

type BackupClient struct {
	EncryptionKey []byte
}

func (bc *BackupClient) Hash(chunk []byte) (hash string) {
	chunkhash := sha256.Sum256(chunk)
	return fmt.Sprintf("sha256:%x", chunkhash)
}

func (bc *BackupClient) PutChunk(chunk []byte) (hash string, err error) {
	hash = bc.Hash(chunk)
	fmt.Printf("Put chunk %s (size %d bytes)\n", hash, len(chunk))
	return hash, nil
}
