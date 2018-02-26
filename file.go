package daobackup

import (
	"os"
	"path/filepath"
)

const (
	TYPE_FILE     = iota
	TYPE_DIR      = iota
	TYPE_SYMLINK  = iota
	TYPE_HARDLINK = iota
)

type FileMeta struct {
	filehandle *os.File
	FullPath   string
	Type       int
	User       string
	Group      string
}

func BackupFile(filename string) (fm *FileMeta, err error) {
	fm = new(FileMeta)
	fm.FullPath, err = filepath.Abs(filename)
	if err != nil {
		return
	}
	fm.filehandle, err = os.Open(fm.FullPath)
	if err != nil {
		return nil, err
	}
	info, err := fm.filehandle.Stat()
	if err != nil {
		return nil, err
	}
	mode := info.Mode()
	switch {
	case mode.IsDir():
		fm.Type = TYPE_DIR
	}
	return fm, err
}
