package daobackup

import (
	"io/ioutil"
	"os"
	"path"
)

type ChunkStoreDisk struct {
	// Location on disk where chunks are stored
	StoragePath string
}

func (csd *ChunkStoreDisk) hashToPath(hash string) string {
	return path.Join(csd.StoragePath, hash[0:2], hash[2:4], hash[4:])
}

func (csd *ChunkStoreDisk) GetChunk(hash string) ([]byte, error) {
	return ioutil.ReadFile(csd.hashToPath(hash))
}

func (csd *ChunkStoreDisk) PutChunk(hash string, data []byte) error {
	if csd.ChunkExists(hash) {
		// Don't try to store chunks that allready exists
		return nil
	}
	fullpath := csd.hashToPath(hash)
	dir := path.Dir(fullpath)
	if err := os.MkdirAll(dir, os.ModeDir|0750); err != nil {
		return err
	}
	return ioutil.WriteFile(fullpath, data, 0640)
}

func (csd *ChunkStoreDisk) ChunkExists(hash string) bool {
	if _, err := os.Stat(csd.hashToPath(hash)); err == nil {
		return true
	}
	return false
}
