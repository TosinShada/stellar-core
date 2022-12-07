package resourceadapter

import (
	"context"
	"fmt"

	"github.com/TosinShada/stellar-core/amount"
	protocol "github.com/TosinShada/stellar-core/protocols/horizon"
	horizonContext "github.com/TosinShada/stellar-core/services/horizon/internal/context"
	"github.com/TosinShada/stellar-core/services/horizon/internal/db2/history"
	"github.com/TosinShada/stellar-core/support/render/hal"
	"github.com/TosinShada/stellar-core/xdr"
)

// PopulateClaimableBalance fills out the resource's fields
func PopulateClaimableBalance(
	ctx context.Context,
	dest *protocol.ClaimableBalance,
	claimableBalance history.ClaimableBalance,
	ledger *history.Ledger,
) error {
	dest.BalanceID = claimableBalance.BalanceID
	dest.Asset = claimableBalance.Asset.StringCanonical()
	dest.Amount = amount.StringFromInt64(int64(claimableBalance.Amount))
	if claimableBalance.Sponsor.Valid {
		dest.Sponsor = claimableBalance.Sponsor.String
	}
	dest.LastModifiedLedger = claimableBalance.LastModifiedLedger
	dest.Claimants = make([]protocol.Claimant, len(claimableBalance.Claimants))
	for i, c := range claimableBalance.Claimants {
		dest.Claimants[i].Destination = c.Destination
		dest.Claimants[i].Predicate = c.Predicate
	}

	if ledger != nil {
		dest.LastModifiedTime = &ledger.ClosedAt
	}

	if xdr.ClaimableBalanceFlags(claimableBalance.Flags).IsClawbackEnabled() {
		dest.Flags.ClawbackEnabled = xdr.ClaimableBalanceFlags(claimableBalance.Flags).IsClawbackEnabled()
	}

	lb := hal.LinkBuilder{Base: horizonContext.BaseURL(ctx)}
	self := fmt.Sprintf("/claimable_balances/%s", dest.BalanceID)
	dest.Links.Self = lb.Link(self)
	dest.PT = fmt.Sprintf("%d-%s", claimableBalance.LastModifiedLedger, dest.BalanceID)
	dest.Links.Transactions = lb.PagedLink(self, "transactions")
	dest.Links.Operations = lb.PagedLink(self, "operations")
	return nil
}
