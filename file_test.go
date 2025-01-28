package daobackup

import (
	"testing"
)

func TestFileChunker(t *testing.T) {
	filepath := "C:\\Users\\floor\\Downloads\\go1.23.5.windows-amd64.msi"

	for h, c := range ChunkFile(filepath) {
		if c.Error != nil {
			t.Error(c.Error)
		}
		t.Logf("%x, %d\n", h, len(c.Bytes))
	}
}

func TestDirChunker(t *testing.T) {
	filepath := "C:\\Users\\floor\\Downloads"

	for h, c := range ChunkDir(filepath) {
		if c.Error != nil {
			t.Error(c.Error)
		}
		t.Logf("%x, %d\n", h, len(c.Bytes))
	}
}
