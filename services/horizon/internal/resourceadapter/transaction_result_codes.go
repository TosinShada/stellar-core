package resourceadapter

import (
	"context"

	protocol "github.com/TosinShada/stellar-core/protocols/horizon"
	"github.com/TosinShada/stellar-core/services/horizon/internal/txsub"
)

// Populate fills out the details
func PopulateTransactionResultCodes(ctx context.Context,
	transactionHash string,
	dest *protocol.TransactionResultCodes,
	fail *txsub.FailedTransactionError,
) (err error) {

	results, err := fail.TransactionResultCodes(transactionHash)
	if err != nil {
		return
	}
	dest.TransactionCode = results.Code
	dest.InnerTransactionCode = results.InnerCode

	dest.OperationCodes, err = fail.OperationResultCodes()
	if err != nil {
		return
	}

	return
}
