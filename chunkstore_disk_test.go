package daobackup

import (
	"io/ioutil"
	//"os"
	"testing"
)

func TestCreateAndRead(t *testing.T) {
	root, err := ioutil.TempDir("/tmp/", "daobackuptest")
	if err != nil {
		t.Error(err)
	}
	//defer os.RemoveAll(root)

	csd := ChunkStoreDisk{StoragePath: root}
	data := []byte("test")
	err = csd.PutChunk("abcdefghijkl", data)
	if err != nil {
		t.Error(err)
	}
}
