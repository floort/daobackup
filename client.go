package daobackup

import (
	"fmt"
	"path/filepath"
	"time"
)

type BackupClient struct {
}

func (b BackupClient) Backup(root string) (err error) {
	backuproot := BackupRoot{}
	fullpath, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	backuproot.Path = fullpath
	hash := ChunkHash{}
	blob := ErrorBytes{}
	for hash, blob = range ChunkDir(root) {
		if blob.Error != nil {
			return err
		}
	}
	backuproot.RootDir = hash
	backuproot.Time = time.Now()
	bytesbuf := make([]byte, 0)
	bytesbuf, err = backuproot.MarshalMsg(bytesbuf)
	if err != nil {
		return err
	}
	fmt.Println(Hash(bytesbuf))
	return nil
}
