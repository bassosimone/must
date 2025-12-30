// SPDX-License-Identifier: GPL-3.0-or-later

package must_test

import (
	"strings"
	"testing"

	"github.com/bassosimone/must"
	"github.com/stretchr/testify/assert"
)

func TestMustFprintf(t *testing.T) {
	var sb strings.Builder
	must.Fprintf(&sb, "%s\n", "Hello, world")
	assert.Equal(t, "Hello, world\n", sb.String())
}
