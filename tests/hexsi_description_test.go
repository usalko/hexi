package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usalko/hexsi"
	"github.com/usalko/hexsi/ft"
)

func TestGzipShortName(t *testing.T) {
	assert.Equal(t, hexsi.FileTypeShortName(ft.GZIP), "GZIP")
}
