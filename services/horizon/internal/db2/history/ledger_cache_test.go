package history

import (
	"testing"

	"github.com/TosinShada/stellar-core/services/horizon/internal/test"
)

func TestLedgerCache(t *testing.T) {
	tt := test.Start(t)
	tt.Scenario("base")
	defer tt.Finish()
	q := &Q{tt.HorizonSession()}

	var lc LedgerCache
	lc.Queue(2)
	lc.Queue(3)

	err := lc.Load(tt.Ctx, q)

	if tt.Assert.NoError(err) {
		tt.Assert.Contains(lc.Records, int32(2))
		tt.Assert.Contains(lc.Records, int32(3))
		tt.Assert.NotContains(lc.Records, int32(1))
	}
}
