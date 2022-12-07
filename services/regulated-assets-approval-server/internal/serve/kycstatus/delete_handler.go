package kycstatus

import (
	"context"
	"net/http"

	"github.com/TosinShada/stellar-core/services/regulated-assets-approval-server/internal/serve/httperror"
	"github.com/TosinShada/stellar-core/support/errors"
	"github.com/TosinShada/stellar-core/support/http/httpdecode"
	"github.com/TosinShada/stellar-core/support/log"
	"github.com/TosinShada/stellar-core/support/render/httpjson"
	"github.com/jmoiron/sqlx"
)

type DeleteHandler struct {
	DB *sqlx.DB
}

func (h DeleteHandler) validate() error {
	if h.DB == nil {
		return errors.New("database cannot be nil")
	}
	return nil
}

type deleteRequest struct {
	StellarAddress string `path:"stellar_address"`
}

func (h DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := h.validate()
	if err != nil {
		log.Ctx(ctx).Error(errors.Wrap(err, "validating kyc-status DeleteHandler"))
		httperror.InternalServer.Render(w)
		return
	}

	in := deleteRequest{}
	err = httpdecode.Decode(r, &in)
	if err != nil {
		log.Ctx(ctx).Error(errors.Wrap(err, "decoding kyc-status DELETE Request"))
		httperror.BadRequest.Render(w)
		return
	}

	err = h.handle(ctx, in)
	if err != nil {
		httpErr, ok := err.(*httperror.Error)
		if !ok {
			httpErr = httperror.InternalServer
		}
		httpErr.Render(w)
		return
	}

	httpjson.Render(w, httpjson.DefaultResponse, httpjson.JSON)
}

func (h DeleteHandler) handle(ctx context.Context, in deleteRequest) error {
	// Check if deleteRequest StellarAddress value is present.
	if in.StellarAddress == "" {
		return httperror.NewHTTPError(http.StatusBadRequest, "Missing stellar address.")
	}

	var existed bool
	const q = `
		WITH deleted_rows AS (
			DELETE FROM accounts_kyc_status
			WHERE stellar_address = $1
			RETURNING *
		) SELECT EXISTS (
			SELECT * FROM deleted_rows
		)
	`
	err := h.DB.QueryRowContext(ctx, q, in.StellarAddress).Scan(&existed)
	if err != nil {
		return errors.Wrap(err, "querying the database")
	}
	if !existed {
		return httperror.NewHTTPError(http.StatusNotFound, "Not found.")
	}

	return nil
}
