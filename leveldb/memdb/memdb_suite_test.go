package memdb

import (
	"testing"

	"github.cbhq.net/cloud/goleveldb/leveldb/testutil"
)

func TestMemDB(t *testing.T) {
	testutil.RunSuite(t, "MemDB Suite")
}
