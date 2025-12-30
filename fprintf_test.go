// SPDX-License-Identifier: GPL-3.0-or-later

package must_test

import (
	"errors"
	"io"
	"strings"
	"testing"

	"github.com/bassosimone/must"
	"github.com/stretchr/testify/assert"
)

func TestMustFprintfSuccess(t *testing.T) {
	var sb strings.Builder
	must.Fprintf(&sb, "%s\n", "Hello, world")
	assert.Equal(t, "Hello, world\n", sb.String())
}

type flexibleWriter struct {
	WriteFn func(data []byte) (int, error)
}

var _ io.Writer = flexibleWriter{}

// Write implements [io.Writer].
func (f flexibleWriter) Write(data []byte) (int, error) {
	return f.WriteFn(data)
}

func TestMustFprintfFailure(t *testing.T) {
	writer := flexibleWriter{
		WriteFn: func(data []byte) (int, error) {
			return 0, errors.New("mocked error")
		},
	}
	assert.Panics(t, func() {
		must.Fprintf(writer, "%s\n", "Hello, world")
	})
}
