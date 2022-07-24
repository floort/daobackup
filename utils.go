package daobackup


import (
	"crypto/sha256"
	"fmt"
)

func Hash(chunk []byte) (hash string) {
	chunkhash := sha256.Sum256(chunk)
	return fmt.Sprintf("sha256:%x", chunkhash)
}