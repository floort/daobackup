package daobackup

import (
	"testing"
)

func TestClient(t *testing.T) {
	bc := BackupClient{}
	bc.Backup("C:\\Users\\floor\\Downloads\\Besluit + documenten")
}
