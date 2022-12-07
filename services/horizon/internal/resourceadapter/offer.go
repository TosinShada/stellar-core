package resourceadapter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/TosinShada/stellar-core/amount"
	protocol "github.com/TosinShada/stellar-core/protocols/horizon"
	horizonContext "github.com/TosinShada/stellar-core/services/horizon/internal/context"
	"github.com/TosinShada/stellar-core/services/horizon/internal/db2/history"
	"github.com/TosinShada/stellar-core/support/render/hal"
	"github.com/TosinShada/stellar-core/xdr"
)

// PopulateOffer constructs an offer response struct from an offer row extracted from the
// the horizon offers table.
func PopulateOffer(ctx context.Context, dest *protocol.Offer, row history.Offer, ledger *history.Ledger) {
	dest.ID = int64(row.OfferID)
	dest.PT = fmt.Sprintf("%d", row.OfferID)
	dest.Seller = row.SellerID
	dest.Amount = amount.String(xdr.Int64(row.Amount))
	dest.PriceR.N = row.Pricen
	dest.PriceR.D = row.Priced
	dest.Price = big.NewRat(int64(row.Pricen), int64(row.Priced)).FloatString(7)
	if row.Sponsor.Valid {
		dest.Sponsor = row.Sponsor.String
	}

	row.SellingAsset.MustExtract(&dest.Selling.Type, &dest.Selling.Code, &dest.Selling.Issuer)
	row.BuyingAsset.MustExtract(&dest.Buying.Type, &dest.Buying.Code, &dest.Buying.Issuer)

	dest.LastModifiedLedger = int32(row.LastModifiedLedger)
	if ledger != nil {
		dest.LastModifiedTime = &ledger.ClosedAt
	}
	lb := hal.LinkBuilder{horizonContext.BaseURL(ctx)}
	dest.Links.Self = lb.Linkf("/offers/%d", row.OfferID)
	dest.Links.OfferMaker = lb.Linkf("/accounts/%s", row.SellerID)
}
