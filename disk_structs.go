package daobackup

import "time"

//go:generate msgp

type ChunkHash [32]byte

type ChunkMeta struct {
	Offset uint64    `msg:"offset"`
	Size   uint64    `msg:"size"`
	Hash   ChunkHash `msg:"hash"`
}

type FileType byte

const (
	Normal FileType = iota
	Directory
	SymbolicLink
	HardLink
)

type FileMeta struct {
	Type       FileType             `msg:"filetype"`
	Hash       ChunkHash            `msg:"hash"`
	Size       int64                `msg:"size"`
	Name       string               `msg:"name"`
	Owner      string               `msg:"owner"`
	Group      string               `msg:"group"`
	Mode       uint32               `msg:"mode"`
	ModTime    time.Time            `msg:"modtime"`
	CreateTime time.Time            `msg:"createtime"`
	Chunks     []ChunkMeta          `msg:"chunks"`
	Entries    map[string]ChunkHash `msg:"entries"`
}

type BackupRoot struct {
	Previous ChunkHash `msg:"previous"`
	Path     string    `msg:"path"`
	Time     time.Time `msg:"time"`
	RootDir  ChunkHash `msg:"rootdir"`
}
