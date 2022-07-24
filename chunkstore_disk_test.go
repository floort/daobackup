package daobackup

import (
	"io/ioutil"
	"os"
	"testing"
	"bytes"
)


func TestCreateAndRead(t *testing.T) {
	root, err := ioutil.TempDir("", "daobackuptest")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(root)

	csd := ChunkStoreDisk{StoragePath: root}
	data := []byte("test")
	data_hash := Hash(data)
	err = csd.PutChunk(data_hash, data)
	if err != nil {
		t.Error(err)
	}
	
	data_read, err := csd.GetChunk(data_hash)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(data, data_read) {
		 t.Fatalf("Data read is not the same as data read.")
	}
	
}




func TestCreateTwice(t *testing.T) {
	root, err := ioutil.TempDir("", "daobackuptest")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(root)

	csd := ChunkStoreDisk{StoragePath: root}
	data := []byte("test")
	data_hash := Hash(data)
	err = csd.PutChunk(data_hash, data)
	if err != nil {
		t.Error(err)
	}
	
	err = csd.PutChunk(data_hash, data)
	if err != nil {
		t.Fatalf("Storing the same chunk twice should not return an error.")
	}
	
}