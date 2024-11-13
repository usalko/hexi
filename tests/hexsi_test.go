package tests

import (
	"encoding/hex"
	"fmt"
	"strings"
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

func wikiBytes(bytesString string) []byte {
	result := []byte{}
	for _, byteAsString := range strings.Split(strings.Trim(bytesString, " "), " ") {
		b, err := hex.DecodeString(byteAsString)
		check(err, "Invalid hex representation of byte %s", byteAsString)
		result = append(result, b[0])
	}
	return result
}

func TestLookupShebang(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 21"))
	check(err)

	assert.Equal(t, *fileType, internal.SHEBANG)
}

func TestLookupClarisWork(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("02 00 5a 57 52 54 00 00 00 00 00 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.CLARIS_WORKS)
}

func TestLookupLotus123V1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 02 00 06 04 06 00 08 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V1)
}

func TestLookupLotus123V3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 1A 00 00 10 04 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V3)
}

func TestLookupLotus123V4V5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 1A 00 02 10 04 00 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V4_V5)
}

func TestLookupLotus123V9(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 1A 00 05 10 04"))
	check(err)

	assert.Equal(t, *fileType, internal.LOTUS_123_V9)
}

func TestLookupAmigaHunkExe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 03 F3"))
	check(err)

	assert.Equal(t, *fileType, internal.AMIGA_HUNK_EXE)
}

func TestLookupQuarkExpress1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 49 49 58 50 52"))
	check(err)

	assert.Equal(t, *fileType, internal.QUARK_EXPRESS)
}

func TestLookupQuarkExpress2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 4D 4D 58 50 52"))
	check(err)

	assert.Equal(t, *fileType, internal.QUARK_EXPRESS)
}

func TestLookupPasswordGorilla(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 57 53 33"))
	check(err)

	assert.Equal(t, *fileType, internal.PASSWORD_GORILLA)
}

func TestLookupLibpcap1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("D4 C3 B2 A1"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP)
}

func TestLookupLibpcap2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("A1 B2 C3 D4"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP)
}

func TestLookupLibpcapNs1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 3C B2 A1"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP_NS)
}

func TestLookupLibpcapNs2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("A1 B2 3C 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.LIBPCAP_NS)
}

func TestLookupPcapng(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0A 0D 0D 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PCAPNPG)
}

func TestLookupRpm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("ED AB EE DB"))
	check(err)

	assert.Equal(t, *fileType, internal.RPM)
}

func TestLookupSqlite3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 51 4C 69 74 65 20 66 6F 72 6D 61 74 20 33 00"))
	check(err)

	assert.Equal(t, *fileType, internal.SQLITE3)
}

func TestLookupAmazonKindleUp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 50 30 31"))
	check(err)

	assert.Equal(t, *fileType, internal.AMAZON_KINDLE_UP)
}

func TestLookupDoomWad(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 57 41 44"))
	check(err)

	assert.Equal(t, *fileType, internal.DOOM_WAD)
}

func TestLookupZero(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00"))
	check(err)

	assert.Equal(t, *fileType, internal.ZERO)
}

// func TestLookupPalmPilot(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("01 01 01 01 01 01 01 01 01 01 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.PALM_PILOT)
// }

func TestLookupPalmDskCalendar(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("BE BA FE CA"))
	check(err)

	assert.Equal(t, *fileType, internal.PALM_DSK_CALENDAR)
}

func TestLookupPalmDskTodo(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 01 42 44"))
	check(err)

	assert.Equal(t, *fileType, internal.PALM_DSK_TODO)
}

func TestLookupPalmDskCalendar2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 01 44 54"))
	check(err)

	assert.Equal(t, *fileType, internal.PALM_DSK_CALENDAR2)
}

func TestLookupTelegramDesktop(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("54 44 46 24"))
	check(err)

	assert.Equal(t, *fileType, internal.TELEGRAM_DSK)
}

func TestLookupTelegramDesktopEncrypted(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("54 44 45 46"))
	check(err)

	assert.Equal(t, *fileType, internal.TELEGRAM_DSK_ENC)
}

func TestLookupPalmDesktopData(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 01 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.PALM_DSK_DATA)
}

func TestLookupIcon(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 01 00"))
	check(err)

	assert.Equal(t, *fileType, internal.ICON)
}

func TestLookupAppleIconFormat(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("69 63 6e 73"))
	check(err)

	assert.Equal(t, *fileType, internal.APPLE_ICON_FORMAT)
}

// func TestLookupTreeGpp(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("66 74 79 70 33 67"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.THREE_GPP)
// }

// func TestLookupHeic(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("66 74 79 70 68 65 69 63 66 74 79 70 6d"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.HEIC)
// }

func TestLookupZlzw(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("1F 9D"))
	check(err)

	assert.Equal(t, *fileType, internal.Z_LZW)
}

func TestLookupZlzh(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("1F A0"))
	check(err)

	assert.Equal(t, *fileType, internal.Z_LZH)
}

// func TestLookupLzh0(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("2D 68 6C 30 2D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.LZH0)
// }

// func TestLookupLzh5(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("2D 68 6C 35 2D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.LZH5)
// }

func TestLookupAmiBack(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("42 41 43 4B 4D 49 4B 45 44 49 53 4B"))
	check(err)

	assert.Equal(t, *fileType, internal.AMI_BACK)
}

func TestLookupAmiBackIdx(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 4E 44 58"))
	check(err)

	assert.Equal(t, *fileType, internal.AMI_BACK_IDX)
}

func TestLookupBplist(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("62 70 6C 69 73 74"))
	check(err)

	assert.Equal(t, *fileType, internal.BPLIST)
}

func TestLookupBz2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("42 5A 68"))
	check(err)

	assert.Equal(t, *fileType, internal.BZ2)
}
