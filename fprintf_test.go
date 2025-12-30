// SPDX-License-Identifier: GPL-3.0-or-later

package must_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/bassosimone/iotest"
	"github.com/bassosimone/must"
	"github.com/stretchr/testify/assert"
)

func TestMustFprintfSuccess(t *testing.T) {
	var sb strings.Builder
	must.Fprintf(&sb, "%s\n", "Hello, world")
	assert.Equal(t, "Hello, world\n", sb.String())
}

func TestMustFprintfFailure(t *testing.T) {
	writer := &iotest.FuncWriter{
		WriteFunc: func(data []byte) (int, error) {
			return 0, errors.New("mocked error")
		},
	}
	assert.Panics(t, func() {
		must.Fprintf(writer, "%s\n", "Hello, world")
	})
}
