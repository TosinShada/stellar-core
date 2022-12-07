package ticker

import (
	"github.com/TosinShada/stellar-core/services/ticker/internal/gql"
	"github.com/TosinShada/stellar-core/services/ticker/internal/tickerdb"
	hlog "github.com/TosinShada/stellar-core/support/log"
)

func StartGraphQLServer(s *tickerdb.TickerSession, l *hlog.Entry, port string) {
	graphql := gql.New(s, l)

	graphql.Serve(port)
}
