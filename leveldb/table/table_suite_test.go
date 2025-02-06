package table

import (
	"testing"

	"github.cbhq.net/cloud/goleveldb/leveldb/testutil"
)

func TestTable(t *testing.T) {
	testutil.RunSuite(t, "Table Suite")
}
