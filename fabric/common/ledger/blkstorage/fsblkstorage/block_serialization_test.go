package fsblkstorage

import ("testing"
	"github.com/hyperledger2/fabric/common/ledger/testutil"
	"fmt"
)

func TestBlockSerialization(t *testing.T) {
	block := testutil.ConstructTestBlock(t, 1, 10, 100)
	fmt.Println(block)
}