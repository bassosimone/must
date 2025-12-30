//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/ooni/probe-cli/blob/v3.20.1/internal/must/must.go
//

package must

import (
	"fmt"
	"io"

	"github.com/bassosimone/runtimex"
)

// Fprintf is like [fmt.Fprintf] but panics in case of failure.
func Fprintf(w io.Writer, format string, args ...any) {
	_ = runtimex.PanicOnError1(fmt.Fprintf(w, format, args...))
}
