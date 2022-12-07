package processors

import (
	logpkg "github.com/TosinShada/stellar-core/support/log"
	"github.com/TosinShada/stellar-core/xdr"
	"github.com/guregu/null"
)

var log = logpkg.DefaultLogger.WithField("service", "ingest")

const maxBatchSize = 100000

func ledgerEntrySponsorToNullString(entry xdr.LedgerEntry) null.String {
	sponsoringID := entry.SponsoringID()

	var sponsor null.String
	if sponsoringID != nil {
		sponsor.SetValid((*sponsoringID).Address())
	}

	return sponsor
}
