// Copyright (c) 2016 The ohmcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package cpuminer

import (
	"github.com/ohmcsuite/ohmclog"
)

// log is a logger that is initialized with no output filters.  This
// means the package will not perform any logging by default until the caller
// requests it.
var log ohmclog.Logger

// The default amount of logging is none.
func init() {
	DisableLog()
}

// DisableLog disables all library log output.  Logging output is disabled
// by default until UseLogger is called.
func DisableLog() {
	log = ohmclog.Disabled
}

// UseLogger uses a specified Logger to output package logging info.
func UseLogger(logger ohmclog.Logger) {
	log = logger
}
