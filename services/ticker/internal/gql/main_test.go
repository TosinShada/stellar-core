package gql

import (
	"testing"

	"github.com/TosinShada/stellar-core/services/ticker/internal/gql/static"
	"github.com/graph-gophers/graphql-go"
)

func TestValidateSchema(t *testing.T) {
	r := resolver{}
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	graphql.MustParseSchema(static.Schema(), &r, opts...)
}
