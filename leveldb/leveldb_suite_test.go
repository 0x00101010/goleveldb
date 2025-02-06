package leveldb

import (
	"testing"

	"github.cbhq.net/cloud/goleveldb/leveldb/testutil"
)

func TestLevelDB(t *testing.T) {
	testutil.RunSuite(t, "LevelDB Suite")
}
