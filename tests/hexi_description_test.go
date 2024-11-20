package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usalko/hexi"
	"github.com/usalko/hexi/ft"
)

func TestGzipShortName(t *testing.T) {
	assert.Equal(t, hexi.FileTypeShortName(ft.GZIP), "GZIP")
}
