package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/floort/daobackup"
)

type FileSystemStorage struct {
	Root          string
	VerifyOnRead  bool
	VerifyOnWrite bool
}

func (s FileSystemStorage) isValidHash(hash string) bool {
	return daobackup.IsValidHash(hash)
}

func (s FileSystemStorage) getPath(hash string) (p string, err error) {
	if !s.isValidHash(hash) {
		return "", fmt.Errorf("'%s' is not a valid hash", hash)
	}
	return filepath.Join(s.Root, hash[0:2], hash[2:4], hash[4:]), nil
}

func (s FileSystemStorage) Put(hash string, blob []byte) error {
	path, err := s.getPath(hash)
	if err != nil {
		return err
	}
	if s.VerifyOnWrite {
		if hash != daobackup.Hash(blob).String() {
			return fmt.Errorf("Expected hash '%s', got '%s'", hash, daobackup.Hash(blob).String())
		}
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		if err = os.MkdirAll(filepath.Dir(path), 0700); err != nil {
			return err
		}
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(blob)
	return nil
}

func (s FileSystemStorage) Exists(hash string) bool {
	path, err := s.getPath(hash)
	if err != nil {
		return false
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (s FileSystemStorage) Get(hash string) (blob []byte, err error) {
	path, err := s.getPath(hash)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	blob, err = io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return blob, nil
}
