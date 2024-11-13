package tests

import (
	"encoding/hex"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
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

func wikiBytes(bytesString string) []byte {
	result := []byte{}
	for _, byteAsString := range strings.Split(strings.Trim(bytesString, " "), " ") {
		b, err := hex.DecodeString(byteAsString)
		check(err, "Invalid hex representation of byte %s", byteAsString)
		result = append(result, b[0])
	}
	return result
}

func DetectFileType(header []byte) (*internal.FileType, error) {
	hexSignature1, _ := internal.LookupSignatureByBytes1(header)
	if hexSignature1 != nil {
		return &hexSignature1.Tag, nil
	}
	hexSignature2, _ := internal.LookupSignatureByBytes2(header)
	if hexSignature2 != nil {
		return &hexSignature2.Tag, nil
	}
	hexSignature3, _ := internal.LookupSignatureByBytes3(header)
	if hexSignature3 != nil {
		return &hexSignature3.Tag, nil
	}
	return nil, fmt.Errorf("unknown file type")
}

func TestLookupShebang(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("23 21"))
	check(err)

	assert.Equal(t, *fileType, internal.SHEBANG)
}

func TestLookupClarisWork(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("02 00 5a 57 52 54 00 00 00 00 00 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.CLARIS_WORKS)
}

func TestLookupLotus123V1(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("00 00 02 00 06 04 06 00 08 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V1)
}

func TestLookupLotus123V3(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("00 00 1A 00 00 10 04 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V3)
}

func TestLookupLotus123V4V5(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("00 00 1A 00 02 10 04 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V4_V5)
}

func TestLookupLotus123V9(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("00 00 1A 00 05 10 04"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V9)
}

func TestLookupAmigaHunkExe(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("00 00 03 F3"))
	check(err)

	assert.Equal(t, *fileType, internal.AMIGA_HUNK_EXE)
}

func TestLookupQuarkExpress1(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("00 00 49 49 58 50 52"))
	check(err)

	assert.Equal(t, *fileType, internal.QUARK_EXPRESS)
}

func TestLookupQuarkExpress2(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("00 00 4D 4D 58 50 52"))
	check(err)

	assert.Equal(t, *fileType, internal.QUARK_EXPRESS)
}

func TestLookupPasswordGorilla(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("50 57 53 33"))
	check(err)

	assert.Equal(t, *fileType, internal.PASSWORD_GORILLA)
}

func TestLookupLibpcap1(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("D4 C3 B2 A1"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP)
}

func TestLookupLibpcap2(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("A1 B2 C3 D4"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP)
}

func TestLookupLibpcapNs1(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("4D 3C B2 A1"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP_NS)
}

func TestLookupLibpcapNs2(t *testing.T) {
	fileType, err := DetectFileType(wikiBytes("A1 B2 3C 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP_NS)
}
