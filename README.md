# Golang Functions that Panic on Failure

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/must)](https://pkg.go.dev/github.com/bassosimone/must) [![Build Status](https://github.com/bassosimone/must/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/must/actions) [![codecov](https://codecov.io/gh/bassosimone/must/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/must)

The `must` Go package contains stdlib-like functions that `panic` on failure.

Typically, you use functions in the `must` package when:

1. they are not expected to fail;

2. failure indicates a programmer error or an unrecoverable condition.

For example:

```Go
import (
	"os"

	"github.com/bassosimone/must"
)

// If we cannot write the stdout we have very serious issues.
must.Fprintf(os.Stdout, "%s\n", "Hello, world!\n")
```

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/must
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Adapted from [ooni/probe-cli](https://github.com/ooni/probe-cli/blob/v3.20.1/internal/must/must.go).
