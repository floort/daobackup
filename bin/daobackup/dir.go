package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

type DirMeta struct {
	Type    string
	Name    string
	Owner   string
	Group   string
	Entries map[string]string
	Errors  map[string]string
}

func (fm *DirMeta) String() string {
	b, err := json.MarshalIndent(fm, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func BackupDir(bc *BackupClient, pathname string) (hash string, err error) {
	fmt.Println(pathname)
	f, err := os.Open(pathname)
	if err != nil {
		return "", err
	}
	defer f.Close()
	stat, err := f.Stat()
	if err != nil {
		return "", err
	}
	if stat.IsDir() == false {
		return BackupFile(bc, pathname)
	}
	dirmeta := DirMeta{
		Type:    "dir",
		Name:    stat.Name(),
		Entries: make(map[string]string),
		Errors:  make(map[string]string),
	}
	entries, err := f.Readdirnames(0)
	if err != nil {
		return "", err
	}
	for e := range entries {
		newpath := path.Join(pathname, entries[e])
		entryhash, err := BackupDir(bc, newpath)
		if err != nil {
			dirmeta.Errors[entries[e]] = err.Error()
		}
		dirmeta.Entries[entries[e]] = entryhash
	}
	return bc.PutChunk([]byte(dirmeta.String()))
}
