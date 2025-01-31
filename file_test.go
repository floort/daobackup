package daobackup

import (
	"testing"
)

func TestFileChunker(t *testing.T) {
	filepath := "C:\\Users\\floor\\Downloads\\go1.23.5.windows-amd64.msi"
	c := ErrorBytes{}
	for _, c = range ChunkFile(filepath) {
		if c.Error != nil {
			t.Error(c.Error)
		}
	}
}

func TestDirChunker(t *testing.T) {
	filepath := "C:\\Users\\floor\\Downloads\\Besluit + documenten"
	c := ErrorBytes{}
	for _, c = range ChunkDir(filepath) {
		if c.Error != nil {
			t.Error(c.Error)
		}
	}
}
