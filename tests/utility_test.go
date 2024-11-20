package tests

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usalko/hexi"
	"github.com/usalko/hexi/ft"
)

func TestSliceMutation(t *testing.T) {
	header := []byte{0x1F, 0x8B, 0x0B, 0x0B}
	initialCapacity := cap(header)
	initialLength := len(header)

	fileType, err := hexi.DetectFileType(header)
	check(err)

	assert.Equal(t, *fileType, ft.GZIP)

	assert.Equal(t, initialCapacity, cap(header))
	assert.Equal(t, initialLength, len(header))
}

func TestNotEmptyStrings(t *testing.T) {
	strings := []string{"", "Hex", "Si"}
	assert.True(t,
		slices.Equal(
			slices.Collect(ft.NotEmptyStrings(strings)),
			[]string{"Hex", "Si"},
		),
	)
}
