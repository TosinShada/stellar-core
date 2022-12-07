package actions

import "github.com/TosinShada/stellar-core/services/horizon/internal/corestate"

type CoreStateGetter interface {
	GetCoreState() corestate.State
}
