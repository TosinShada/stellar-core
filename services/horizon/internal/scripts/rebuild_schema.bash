#! /usr/bin/env bash
set -e

# This scripts rebuilds the latest.sql file included in the schema package.
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
GOTOP="$( cd "$DIR/../../../../../../../.." && pwd )"

go generate github.com/TosinShada/stellar-core/services/horizon/internal/db2/schema
go generate github.com/TosinShada/stellar-core/services/horizon/internal/test
go install github.com/TosinShada/stellar-core/services/horizon
