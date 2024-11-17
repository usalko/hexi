package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usalko/hexsi"
	"github.com/usalko/hexsi/ft"
)

func TestSliceMutation(t *testing.T) {
	header := []byte{0x1F, 0x8B, 0x0B, 0x0B}
	initialCapacity := cap(header)
	initialLength := len(header)

	fileType, err := hexsi.DetectFileType(header)
	check(err)

	assert.Equal(t, *fileType, ft.GZIP)

	assert.Equal(t, initialCapacity, cap(header))
	assert.Equal(t, initialLength, len(header))
}
