package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/usalko/hexsi"
	"github.com/usalko/hexsi/internal"
)

func check(err error, msgs ...any) {
	if err != nil {
		if len(msgs) == 0 {
			panic(err)
		} else if len(msgs) == 1 {
			panic(fmt.Errorf("%s: %s", msgs[0], err))
		} else {
			panic(fmt.Errorf("%s: %s", fmt.Sprintf(msgs[0].(string), msgs[1:]...), err))
		}
	}
}

func TestStreamReader(t *testing.T) {
	fileType, err := hexsi.DetectFileType([]byte{0x23, 0x21})
	check(err)

	assert.Equal(t, *fileType, internal.SHEBANG)
}
