package daobackup

import (
	"fmt"
	"path/filepath"

	"google.golang.org/protobuf/proto"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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
	backuproot.RootDirectory = hash.Bytes()
	backuproot.Timestamp = timestamppb.Now()
	bytesbuf, err := proto.Marshal(&backuproot)
	if err != nil {
		return err
	}
	fmt.Println(HashChunk(bytesbuf))
	return nil
}
