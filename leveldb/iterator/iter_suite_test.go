package iterator_test

import (
	"testing"

	"github.cbhq.net/cloud/goleveldb/leveldb/testutil"
)

func TestIterator(t *testing.T) {
	testutil.RunSuite(t, "Iterator Suite")
}
