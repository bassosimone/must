//
// SPDX-License-Identifier: GPL-3.0-or-later
//
// Adapted from: https://github.com/ooni/probe-cli/blob/v3.20.1/internal/must/must.go
//

// Package must contains stdlib-like functions that panic on error.
//
// Use this package when you know that the function cannot fail or when the
// failure indicates a bug or unforeseen error, for which a stack trace will
// be more useful than just an error message.
package must
