package test

import (
	"github.com/TosinShada/stellar-core/support/log"
	"github.com/sirupsen/logrus"
)

var testLogger *log.Entry

func init() {
	testLogger = log.New()
	testLogger.DisableColors()
	testLogger.SetLevel(logrus.DebugLevel)
}
