package daobackup

import (
	context "context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type BasicFilesystemServer struct {
	Root          string
	VerifyOnRead  bool
	VerifyOnWrite bool
}

func (server BasicFilesystemServer) getPath(hash ChunkHash) string {
	hashstring := hash.String()
	return filepath.Join(server.Root, hashstring[0:2], hashstring[2:4], hashstring[4:]) // /root/path/aa/bb/xxxxxxxxxx for hash aabbxxxxxxxxxx
}

func (server *BasicFilesystemServer) PutBlob(c context.Context, hashedblob *HashedBlob) (*RPCStatus, error) {
	blobhash := ChunkHashFromBytes(hashedblob.Hash)
	path := server.getPath(blobhash)
	if server.VerifyOnWrite {
		if !blobhash.Verify(hashedblob.Blob) {
			err := fmt.Errorf("expected hash '%s', got different hash when hashing provided blob", blobhash)
			return &RPCStatus{Succes: false, Message: err.Error()}, err
		}
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			return &RPCStatus{Succes: false, Message: err.Error()}, err
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return &RPCStatus{Succes: false, Message: err.Error()}, err
	}
	defer f.Close()
	f.Write(hashedblob.Blob)
	return &RPCStatus{Succes: true, Message: ""}, nil
}
func (server *BasicFilesystemServer) CheckBlob(c context.Context, chunkhash *Hash) (*RPCStatus, error) {
	ch := ChunkHashFromBytes(chunkhash.Hash)
	path := server.getPath(ch)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return &RPCStatus{Succes: false, Message: ""}, nil
	}
	return &RPCStatus{Succes: true, Message: ""}, nil
}

func (server *BasicFilesystemServer) GetBlob(c context.Context, chunkhash *Hash) (*HashedBlob, error) {
	ch := ChunkHashFromBytes(chunkhash.Hash)
	path := server.getPath(ch)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	blob, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	hb := &HashedBlob{Blob: blob}
	if server.VerifyOnRead {
		hb.Hash = HashChunk(blob).Bytes()
	} else {
		hb.Hash = ch.Bytes()
	}
	return hb, nil
}

func (server *BasicFilesystemServer) mustEmbedUnimplementedDAOBackupServer() {}
