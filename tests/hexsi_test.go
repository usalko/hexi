package tests

import (
	"crypto/rand"
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
		if byteAsString == "??" {
			b := make([]byte, 1)
			if _, err := rand.Read(b); err != nil {
				check(err)
			}
			result = append(result, b[0])
		} else {
			b, err := hex.DecodeString(byteAsString)
			check(err, "Invalid hex representation of byte %s", byteAsString)
			result = append(result, b[0])
		}
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

// func TestLookupPalmPilot(t *testing.T) { // Offset 11
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

// func TestLookupTreeGpp(t *testing.T) { // Offset 4
// 	fileType, err := hexsi.DetectFileType(wikiBytes("66 74 79 70 33 67"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.THREE_GPP)
// }

// func TestLookupHeic(t *testing.T) { // Offset 4
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

// func TestLookupLzh0(t *testing.T) { // Offset 2
// 	fileType, err := hexsi.DetectFileType(wikiBytes("2D 68 6C 30 2D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.LZH0)
// }

// func TestLookupLzh5(t *testing.T) { // Offset 2
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

func TestLookupGif87a(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("47 49 46 38 37 61"))
	check(err)

	assert.Equal(t, *fileType, internal.GIF)
}

func TestLookupGif89a(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("47 49 46 38 39 61"))
	check(err)

	assert.Equal(t, *fileType, internal.GIF)
}

func TestLookupTiffLe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 49 2A 00"))
	check(err)

	assert.Equal(t, *fileType, internal.TIFF)
}

func TestLookupTiffBe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 4D 00 2A"))
	check(err)

	assert.Equal(t, *fileType, internal.TIFF)
}

func TestLookupBigTiffLe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 49 2B 00"))
	check(err)

	assert.Equal(t, *fileType, internal.BIG_TIFF)
}

func TestLookupBigTiffBe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 4D 00 2B"))
	check(err)

	assert.Equal(t, *fileType, internal.BIG_TIFF)
}

func TestLookupKodakCin(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("80 2A 5F D7"))
	check(err)

	assert.Equal(t, *fileType, internal.KODAK_CIN)
}

func TestLookupKodakRncV1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 4E 43 01"))
	check(err)

	assert.Equal(t, *fileType, internal.RNC_V1)

}

func TestLookupKodakRncV2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 4E 43 02"))
	check(err)

	assert.Equal(t, *fileType, internal.RNC_V2)

}

func TestLookupNuruImage(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4E 55 52 55 49 4D 47"))
	check(err)

	assert.Equal(t, *fileType, internal.NURU_IMAGE)

}

func TestLookupNuruPalette(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4E 55 52 55 50 41 4C"))
	check(err)

	assert.Equal(t, *fileType, internal.NURU_PALETTE)

}

func TestLookupSmpteDpxBe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 44 50 58"))
	check(err)

	assert.Equal(t, *fileType, internal.SMPTE_DPX)

}

func TestLookupSmpteDpxLe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("58 50 44 53"))
	check(err)

	assert.Equal(t, *fileType, internal.SMPTE_DPX)

}

func TestLookupOpenExr(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("76 2F 31 01"))
	check(err)

	assert.Equal(t, *fileType, internal.OPEN_EXR)

}

func TestLookupBpg(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("42 50 47 FB"))
	check(err)

	assert.Equal(t, *fileType, internal.BPG)

}

func TestLookupJpegRaw1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF D8 FF DB"))
	check(err)

	assert.Equal(t, *fileType, internal.JPEG_RAW)

}

func TestLookupJpegRaw2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF D8 FF E0 00 10 4A 46 49 46 00 01"))
	check(err)

	assert.Equal(t, *fileType, internal.JPEG_RAW)

}

func TestLookupJpegRaw3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF D8 FF EE"))
	check(err)

	assert.Equal(t, *fileType, internal.JPEG_RAW)

}

func TestLookupJpegRaw4(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF D8 FF E1 ?? ?? 45 78 69 66 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.JPEG_RAW)

}

func TestLookupJpegRaw5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF D8 FF E0"))
	check(err)

	assert.Equal(t, *fileType, internal.JPEG_RAW)

}

func TestLookupJpeg2000Case1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 00 0C 6A 50 20 20 0D 0A 87 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.JPEG_2000)

}

func TestLookupJpeg2000Case2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF 4F FF 51"))
	check(err)

	assert.Equal(t, *fileType, internal.JPEG_2000)

}

func TestLookupQui(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("71 6f 69 66"))
	check(err)

	assert.Equal(t, *fileType, internal.QUI)

}

func TestLookupIifIlbm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 49 4C 42 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_ILBM)

}

func TestLookupIifVoice(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 38 53 56 58"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_VOICE)

}

func TestLookupIifAmigaCb(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 43 42 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_AMIGA_CB)

}

func TestLookupIifAniBmp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 4E 42 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_ANI_BMP)

}

func TestLookupAniCel(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 4E 49 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_ANI_CEL)

}

func TestLookupFaxImg(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 46 41 58 58"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_FAX_IMG)

}

func TestLookupIifFt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 46 54 58 54"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_FT)

}

func TestLookupIifMusScoreSimple(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 53 4D 55 53"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_MUS_SCORE_SIMPLE)

}

func TestLookupIifMusScore(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 43 4D 55 53"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_MUS_SCORE)

}

func TestLookupIifYuvImage(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 59 55 56 4E"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_YUV_IMAGE)

}

func TestLookupAmigaFvm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 46 41 4E 54"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_AMIGA_FVM)

}

func TestLookupIifAiff(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 49 46 46"))
	check(err)

	assert.Equal(t, *fileType, internal.IIF_AIFF)

}

func TestLookupLz(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4C 5A 49 50"))
	check(err)

	assert.Equal(t, *fileType, internal.LZ)

}

func TestLookupCpio(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("30 37 30 37 30 37"))
	check(err)

	assert.Equal(t, *fileType, internal.CPIO)

}

func TestLookupDosMz(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 5A"))
	check(err)

	assert.Equal(t, *fileType, internal.DOS_MZ)

}
func TestLookupSmartSniff(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 4D 53 4E 46 32 30 30"))
	check(err)

	assert.Equal(t, *fileType, internal.SMART_SNIFF)

}
func TestLookupDosZm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("5A 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.DOS_ZM)

}
func TestLookupZipCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 4B 03 04"))
	check(err)

	assert.Equal(t, *fileType, internal.ZIP)

}
func TestLookupZipCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 4B 05 06"))
	check(err)

	assert.Equal(t, *fileType, internal.ZIP)

}
func TestLookupZipCase3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 4B 07 08"))
	check(err)

	assert.Equal(t, *fileType, internal.ZIP)

}
func TestLookupRarV1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 61 72 21 1A 07 00"))
	check(err)

	assert.Equal(t, *fileType, internal.RAR_V1)

}
func TestLookupRarV5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 61 72 21 1A 07 01 00"))
	check(err)

	assert.Equal(t, *fileType, internal.RAR_V5)

}
func TestLookupElf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("7F 45 4C 46"))
	check(err)

	assert.Equal(t, *fileType, internal.ELF)

}
func TestLookupPng(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("89 50 4E 47 0D 0A 1A 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PNG)

}
func TestLookupHdf4(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0E 03 13 01"))
	check(err)

	assert.Equal(t, *fileType, internal.HDF4)

}
func TestLookupHdf5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("89 48 44 46 0D 0A 1A 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.HDF5)

}
func TestLookupCom(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("C9"))
	check(err)

	assert.Equal(t, *fileType, internal.COM)

}
func TestLookupJavaClass(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("CA FE BA BE"))
	check(err)

	assert.Equal(t, *fileType, internal.JAVA_CLASS)

}
func TestLookupUtf8Text(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("EF BB BF"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF8_TXT)

}

func TestLookupUtf16LeTxt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF FE"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF16LE_TXT)

}

func TestLookupUtf16BeTxt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FE FF"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF16BE_TXT)

}
func TestLookupUtf32LeTxt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF FE 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF32LE_TXT)

}
func TestLookupUtf32BeTxt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 FE FF"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF32BE_TXT)

}
func TestLookupUtf7TxtCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2B 2F 76 38"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF7_TXT)

}
func TestLookupUtf7TxtCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2B 2F 76 39"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF7_TXT)

}
func TestLookupUtf7TxtCase3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2B 2F 76 2B"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF7_TXT)

}
func TestLookupUtf7TxtCase4(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2B 2F 76 2F"))
	check(err)

	assert.Equal(t, *fileType, internal.UTF7_TXT)

}
func TestLookupScsuTxt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0E FE FF"))
	check(err)

	assert.Equal(t, *fileType, internal.SCSU_TXT)

}
func TestLookupEbcdicTxt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("DD 73 66 73"))
	check(err)

	assert.Equal(t, *fileType, internal.EBCDIC_TXT)

}
func TestLookupMachOBin32(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FE ED FA CE"))
	check(err)

	assert.Equal(t, *fileType, internal.MACHO_BIN32)

}
func TestLookupMachOBin64(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FE ED FA CF"))
	check(err)

	assert.Equal(t, *fileType, internal.MACHO_BIN64)

}
func TestLookupJks(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FE ED FE ED"))
	check(err)

	assert.Equal(t, *fileType, internal.JKS)

}
func TestLookupMachOBin32R(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("CE FA ED FE"))
	check(err)

	assert.Equal(t, *fileType, internal.MACHO_BIN32R)

}
func TestLookupMachOBin64R(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("CF FA ED FE"))
	check(err)

	assert.Equal(t, *fileType, internal.MACHO_BIN64R)

}
func TestLookupPs(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("25 21 50 53"))
	check(err)

	assert.Equal(t, *fileType, internal.PS)

}
func TestLookupEps3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("25 21 50 53 2D 41 64 6F 62 65 2D 33 2E 30 20 45 50 53 46 2D 33 2E 30"))
	check(err)

	assert.Equal(t, *fileType, internal.EPS3)

}
func TestLookupEps31(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("25 21 50 53 2D 41 64 6F 62 65 2D 33 2E 31 20 45 50 53 46 2D 33 2E 30"))
	check(err)

	assert.Equal(t, *fileType, internal.EPS31)

}
func TestLookupChm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 54 53 46 03 00 00 00 60 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.CHM)

}
func TestLookupHlp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3F 5F"))
	check(err)

	assert.Equal(t, *fileType, internal.HLP)

}
func TestLookupPdf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("25 50 44 46 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.PDF)

}
func TestLookupAsf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("30 26 B2 75 8E 66 CF 11 A6 D9 00 AA 00 62 CE 6C"))
	check(err)

	assert.Equal(t, *fileType, internal.ASF)

}
func TestLookupMsSdi(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("24 53 44 49 30 30 30 31"))
	check(err)

	assert.Equal(t, *fileType, internal.MSSDI)

}
func TestLookupOgg(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4F 67 67 53"))
	check(err)

	assert.Equal(t, *fileType, internal.OGG)

}
func TestLookupPsd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("38 42 50 53"))
	check(err)

	assert.Equal(t, *fileType, internal.PSD)

}

func TestLookupWav(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 57 41 56 45"))
	check(err)

	assert.Equal(t, *fileType, internal.WAV)

}

func TestLookupMp3Case1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF FB"))
	check(err)

	assert.Equal(t, *fileType, internal.MP3)

}
func TestLookupMp3Case2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF F3"))
	check(err)

	assert.Equal(t, *fileType, internal.MP3)

}
func TestLookupMp3Case3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF F2"))
	check(err)

	assert.Equal(t, *fileType, internal.MP3)

}
func TestLookupMp3V2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 44 33"))
	check(err)

	assert.Equal(t, *fileType, internal.MP3V2)

}
func TestLookupBmp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("42 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.BMP)

}

// func TestLookupIso(t *testing.T) { // Offsets: one of 0x8001, 0x8801, 0x9001
// 	fileType, err := hexsi.DetectFileType(wikiBytes("43 44 30 30 31"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.ISO)

// }

// func TestLookupCdi(t *testing.T) {  // Offset 0x5EAC9
// 	fileType, err := hexsi.DetectFileType(wikiBytes("43 44 30 30 31"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.CDI)

// }

func TestLookupMgw(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("6D 61 69 6E 2E 62 73"))
	check(err)

	assert.Equal(t, *fileType, internal.MGW)

}
func TestLookupNes(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4E 45 53"))
	check(err)

	assert.Equal(t, *fileType, internal.NES)

}

// func TestLookupD64(t *testing.T) {  // Offset 0x165A4
// 	fileType, err := hexsi.DetectFileType(wikiBytes("A0 32 41 A0 A0 A0"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.D64)

// }
func TestLookupG64(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("47 53 52 2D 31 35 34 31"))
	check(err)

	assert.Equal(t, *fileType, internal.G64)

}

// func TestLookupD81(t *testing.T) {  // Offset 0x61819
// 	fileType, err := hexsi.DetectFileType(wikiBytes("A0 33 44 A0 A0"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.D81)

// }
func TestLookupT64(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("43 36 34 20 74 61 70 65 20 69 6D 61 67 65 20 66 69 6C 65"))
	check(err)

	assert.Equal(t, *fileType, internal.T64)

}
func TestLookupCrt64(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("43 36 34 20 43 41 52 54 52 49 44 47 45 20 20 20"))
	check(err)

	assert.Equal(t, *fileType, internal.CRT64)

}
func TestLookupFits(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 49 4D 50 4C 45 20 20 3D 20 20 20 20 20 20 20 20 20 20 20 20 20 20 20 20 20 20 20 20 54"))
	check(err)

	assert.Equal(t, *fileType, internal.FITS)

}
func TestLookupFlac(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("66 4C 61 43"))
	check(err)

	assert.Equal(t, *fileType, internal.FLAC)

}
func TestLookupMidi(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 54 68 64"))
	check(err)

	assert.Equal(t, *fileType, internal.MIDI)

}
func TestLookupMsCom(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("D0 CF 11 E0 A1 B1 1A E1"))
	check(err)

	assert.Equal(t, *fileType, internal.MSCOM)

}
func TestLookupDex(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("64 65 78 0A 30 33 35 00"))
	check(err)

	assert.Equal(t, *fileType, internal.DEX)

}
func TestLookupVdmk(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4B 44 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.VDMK)

}

func TestLookupVmware4Vddf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 20 44 69 73 6B 20 44 65 73 63 72 69 70 74 6F"))
	check(err)

	assert.Equal(t, *fileType, internal.VMWARE4VDDF)
}

func TestLookupCrx(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("43 72 32 34"))
	check(err)

	assert.Equal(t, *fileType, internal.CRX)
}

func TestLookupFh8(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("41 47 44 33"))
	check(err)

	assert.Equal(t, *fileType, internal.FH8)
}

func TestLookupAppleWorks5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("05 07 00 00 42 4F 42 4F 05 07 00 00 00 00 00 00 00 00 00 00 00 01"))
	check(err)

	assert.Equal(t, *fileType, internal.APPLE_WORKS5)
}

func TestLookupAppleWorks6(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("06 07 E1 00 42 4F 42 4F 06 07 E1 00 00 00 00 00 00 00 00 00 00 01"))
	check(err)

	assert.Equal(t, *fileType, internal.APPLE_WORKS6)
}

func TestLookupRoxioToastCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("45 52 02 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.ROXIO_TOAST)
}

func TestLookupRoxioToastCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("8B 45 52 02 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.ROXIO_TOAST)
}

// func TestLookupAppleDmg(t *testing.T) {  // Offset from the end-512
// 	fileType, err := hexsi.DetectFileType(wikiBytes("6B 6F 6C 79"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.APPLE_DMG)
// }

func TestLookupXar(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 61 72 21"))
	check(err)

	assert.Equal(t, *fileType, internal.XAR)
}

func TestLookupWindowsDat(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 4D 4F 43 43 4D 4F 43"))
	check(err)

	assert.Equal(t, *fileType, internal.WINDOWS_DAT)
}

func TestLookupNintendoSysRom(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4E 45 53 1A"))
	check(err)

	assert.Equal(t, *fileType, internal.NINTENDO_SYS_ROM)
}

// func TestLookupTarCase1(t *testing.T) {  // Offset 257
// 	fileType, err := hexsi.DetectFileType(wikiBytes("75 73 74 61 72 00 30"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.TAR)
// }

// func TestLookupTarCase2(t *testing.T) {  // Offset 257
// 	fileType, err := hexsi.DetectFileType(wikiBytes("30 75 73 74 61 72 20 20 00"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.TAR)
// }

func TestLookupOar(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4F 41 52 ??"))
	check(err)

	assert.Equal(t, *fileType, internal.OAR)
}

func TestLookupTox(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("74 6F 78 33"))
	check(err)

	assert.Equal(t, *fileType, internal.TOX)
}

func TestLookupMlv(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 4C 56 49"))
	check(err)

	assert.Equal(t, *fileType, internal.MLV)
}

func TestLookupWindowsUpdateCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("44 43 4D 01 50 41 33 30"))
	check(err)

	assert.Equal(t, *fileType, internal.WINDOWS_UPDATE)
}

func TestLookupWindowsUpdateCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 41 33 30"))
	check(err)

	assert.Equal(t, *fileType, internal.WINDOWS_UPDATE)
}

func TestLookupSevenZ(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("37 7A BC AF 27 1C"))
	check(err)

	assert.Equal(t, *fileType, internal.SEVEN_Z)
}

func TestLookupGzip(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("1F 8B"))
	check(err)

	assert.Equal(t, *fileType, internal.GZIP)
}

func TestLookupXz(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FD 37 7A 58 5A 00"))
	check(err)

	assert.Equal(t, *fileType, internal.XZ)
}

func TestLookupLz4(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("04 22 4D 18"))
	check(err)

	assert.Equal(t, *fileType, internal.LZ4)
}

func TestLookupCab(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 53 43 46"))
	check(err)

	assert.Equal(t, *fileType, internal.CAB)
}

func TestLookupMsQuantum(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 5A 44 44 88 F0 27 33"))
	check(err)

	assert.Equal(t, *fileType, internal.MS_QUANTUM)
}

func TestLookupFlif(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4C 49 46"))
	check(err)

	assert.Equal(t, *fileType, internal.FLIF)
}

func TestLookupMkv(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("1A 45 DF A3"))
	check(err)

	assert.Equal(t, *fileType, internal.MKV)
}

func TestLookupStg(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 49 4C 20"))
	check(err)

	assert.Equal(t, *fileType, internal.STG)
}

func TestLookupDjvu(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("41 54 26 54 46 4F 52 4D ?? ?? ?? ?? 44 4A 56"))
	check(err)

	assert.Equal(t, *fileType, internal.DJVU)
}

func TestLookupDer(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("30 82"))
	check(err)

	assert.Equal(t, *fileType, internal.DER)
}

func TestLookupPemCrt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 2D 2D 2D 2D 42 45 47 49 4E 20 43 45 52 54 49 46 49 43 41 54 45 2D 2D 2D 2D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.PEM_CRT)
}

func TestLookupPemCsr(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 2D 2D 2D 2D 42 45 47 49 4E 20 43 45 52 54 49 46 49 43 41 54 45 20 52 45 51 55 45 53 54 2D 2D 2D 2D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.PEM_CSR)
}

func TestLookupPemKeyPkcs8(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 2D 2D 2D 2D 42 45 47 49 4E 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.PEM_KEY_PKCS8)
}

func TestLookupPemKeyDsa(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 2D 2D 2D 2D 42 45 47 49 4E 20 44 53 41 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.PEM_KEY_DSA)
}

func TestLookupPemKeyRsa(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 2D 2D 2D 2D 42 45 47 49 4E 20 52 45 41 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.PEM_KEY_RSA)
}

func TestLookupPpk2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 75 54 54 59 2D 55 73 65 72 2D 4B 65 79 2D 46 69 6C 65 2D 32 3A"))
	check(err)

	assert.Equal(t, *fileType, internal.PPK2)
}

func TestLookupPpk3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 75 54 54 59 2D 55 73 65 72 2D 4B 65 79 2D 46 69 6C 65 2D 33 3A"))
	check(err)

	assert.Equal(t, *fileType, internal.PPK3)
}

func TestLookupOpensshPrivk(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 2D 2D 2D 2D 42 45 47 49 4E 20 4F 50 45 4E 53 53 48 20 50 52 49 56 41 54 45 20 4B 45 59 2D 2D 2D 2D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.OPENSSH_PRIVK)
}

func TestLookupOpensshPubk(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 2D 2D 2D 2D 42 45 47 49 4E 20 53 53 48 32 20 4B 45 59 2D 2D 2D 2D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.OPENSSH_PUBK)
}

// func TestLookupDcm(t *testing.T) { // Offset 128
// 	fileType, err := hexsi.DetectFileType(wikiBytes("44 49 43 4D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.DCM)
// }

func TestLookupWoff(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("77 4F 46 46"))
	check(err)

	assert.Equal(t, *fileType, internal.WOFF)
}

func TestLookupWoff2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("77 4F 46 32"))
	check(err)

	assert.Equal(t, *fileType, internal.WOFF2)
}

func TestLookupXml8(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3C 3F 78 6D 6C 20"))
	check(err)

	assert.Equal(t, *fileType, internal.XML8)
}

func TestLookupXml16Le(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3C 00 3F 00 78 00 6D 00 6C 00 20"))
	check(err)

	assert.Equal(t, *fileType, internal.XML16LE)
}

func TestLookupXml16Be(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 3C 00 3F 00 78 00 6D 00 6C 00 20"))
	check(err)

	assert.Equal(t, *fileType, internal.XML16BE)
}

func TestLookupXml32Le(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3C 00 00 00 3F 00 00 00 78 00 00 00 6D 00 00 00 6C 00 00 00 20 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.XML32LE)
}

func TestLookupXml32Be(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 00 3C 00 00 00 3F 00 00 00 78 00 00 00 6D 00 00 00 6C 00 00 00 20"))
	check(err)

	assert.Equal(t, *fileType, internal.XML32BE)
}

func TestLookupXmlEbcdic(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4C 6F A7 94 93 40"))
	check(err)

	assert.Equal(t, *fileType, internal.XML_EBCDIC)
}

func TestLookupWasm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 61 73 6D"))
	check(err)

	assert.Equal(t, *fileType, internal.WASM)
}

func TestLookupLep(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("CF 84 01"))
	check(err)

	assert.Equal(t, *fileType, internal.LEP)
}

func TestLookupSwfCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("43 57 53"))
	check(err)

	assert.Equal(t, *fileType, internal.SWF)
}

func TestLookupSwfCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 57 53"))
	check(err)

	assert.Equal(t, *fileType, internal.SWF)
}

func TestLookupDeb(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("21 3C 61 72 63 68 3E 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.DEB)
}

func TestLookupWebp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 57 45 42 50"))
	check(err)

	assert.Equal(t, *fileType, internal.WEBP)
}

func TestLookupUboot(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("27 05 19 56"))
	check(err)

	assert.Equal(t, *fileType, internal.UBOOT)
}

func TestLookupRtf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("7B 5C 72 74 66 31"))
	check(err)

	assert.Equal(t, *fileType, internal.RTF)
}

func TestLookupMstape(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("54 41 50 45"))
	check(err)

	assert.Equal(t, *fileType, internal.MSTAPE)
}

func TestLookupMpegTs(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("47"))
	check(err)

	assert.Equal(t, *fileType, internal.MPEG_TS)
}

func TestLookupMpegPs(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 01 BA"))
	check(err)

	assert.Equal(t, *fileType, internal.MPEG_PS)
}

func TestLookupMpegVs(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 01 B3"))
	check(err)

	assert.Equal(t, *fileType, internal.MPEG_VS)
}

// func TestLookupMp4Iso(t *testing.T) { // Offset 4
// 	fileType, err := hexsi.DetectFileType(wikiBytes("66 74 79 70 69 73 6F 6D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.MP4_ISO)
// }

// func TestLookupMp4(t *testing.T) { // Offset 4
// 	fileType, err := hexsi.DetectFileType(wikiBytes("66 74 79 70 4D 53 4E 56"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.MP4)
// }

func TestLookupZlibNc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 01"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_NC)
}

func TestLookupZlibBs(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 5E"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_BS)
}

func TestLookupZlibDc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 9C"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_DC)
}

func TestLookupZlibBc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 DA"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_BC)
}

func TestLookupZlibNcpd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 20"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_NCPD)
}

func TestLookupZlibBspd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 7D"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_BSPD)
}

func TestLookupZlibDcpd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 BB"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_DCPD)
}

func TestLookupZlibBcpd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 F9"))
	check(err)

	assert.Equal(t, *fileType, internal.ZLIB_BCPD)
}

func TestLookupLzfse(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("62 76 78 32"))
	check(err)

	assert.Equal(t, *fileType, internal.LZFSE)
}

func TestLookupOrc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4F 52 43"))
	check(err)

	assert.Equal(t, *fileType, internal.ORC)
}

func TestLookupAvro(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4F 62 6A 01"))
	check(err)

	assert.Equal(t, *fileType, internal.AVRO)
}

func TestLookupRcCff(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 45 51 36"))
	check(err)

	assert.Equal(t, *fileType, internal.RC_CFF)
}

func TestLookupRbxl(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3C 72 6F 62 6C 6F 78 21"))
	check(err)

	assert.Equal(t, *fileType, internal.RBXL)
}

func TestLookupPhotocapOtmpl(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("65 87 78 56"))
	check(err)

	assert.Equal(t, *fileType, internal.PHOTOCAP_OTMPL)
}

func TestLookupPhotocapVec(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("55 55 AA AA"))
	check(err)

	assert.Equal(t, *fileType, internal.PHOTOCAP_VEC)
}

func TestLookupPhotocapTmpl(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("78 56 34"))
	check(err)

	assert.Equal(t, *fileType, internal.PHOTOCAP_TMPL)
}

func TestLookupParquet(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 41 52 31"))
	check(err)

	assert.Equal(t, *fileType, internal.PARQUET)
}

func TestLookupEz2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("45 4D 58 32"))
	check(err)

	assert.Equal(t, *fileType, internal.EZ2)
}

func TestLookupEz3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("45 4D 55 33"))
	check(err)

	assert.Equal(t, *fileType, internal.EZ3)
}

func TestLookupLuaBytecode(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("1B 4C 75 61"))
	check(err)

	assert.Equal(t, *fileType, internal.LUA_BYTECODE)
}

func TestLookupMacosSymlinkCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("62 6F 6F 6B 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.MACOS_SYMLINK)
}

func TestLookupMacosSymlinkCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("6D 61 72 6B 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.MACOS_SYMLINK)
}

func TestLookupMsziCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("5B 5A 6F 6E 65 54 72 61"))
	check(err)

	assert.Equal(t, *fileType, internal.MSZI)
}

func TestLookupMsziCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("6E 73 66 65 72 5D"))
	check(err)

	assert.Equal(t, *fileType, internal.MSZI)
}

func TestLookupEml(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 65 63 65 69 76 65 64 3A"))
	check(err)

	assert.Equal(t, *fileType, internal.EML)
}

func TestLookupTde(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("20 02 01 62 A0 1E AB 07 02 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.TDE)
}

func TestLookupKdb(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("37 48 03 02 00 00 00 00 58 35 30 39 4B 45 59"))
	check(err)

	assert.Equal(t, *fileType, internal.KDB)
}

func TestLookupPgp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("85 ?? ?? 03"))
	check(err)

	assert.Equal(t, *fileType, internal.PGP)
}

func TestLookupZst(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("28 B5 2F FD"))
	check(err)

	assert.Equal(t, *fileType, internal.ZST)
}

func TestLookupQuickZipRs(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 53 56 4B 44 41 54 41"))
	check(err)

	assert.Equal(t, *fileType, internal.QUICK_ZIP_RS)
}

func TestLookupSml(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3A 29 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.SML)
}

func TestLookupPef(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4A 6F 79 21"))
	check(err)

	assert.Equal(t, *fileType, internal.PEF)
}

func TestLookupSrt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("31 0A 30 30"))
	check(err)

	assert.Equal(t, *fileType, internal.SRT)
}

func TestLookupVpk(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("34 12 AA 55"))
	check(err)

	assert.Equal(t, *fileType, internal.VPK)
}

// func TestLookupAce(t *testing.T) {  // Offset 7
// 	fileType, err := hexsi.DetectFileType(wikiBytes("2A 2A 41 43 45 2A 2A"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.ACE)
// }

func TestLookupArj(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("60 EA"))
	check(err)

	assert.Equal(t, *fileType, internal.ARJ)
}

func TestLookupIsCab(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 53 63 28"))
	check(err)

	assert.Equal(t, *fileType, internal.IS_CAB)
}

func TestLookupKwaj(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4B 57 41 4A"))
	check(err)

	assert.Equal(t, *fileType, internal.KWAJ)
}

func TestLookupSzdd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("53 5A 44 44"))
	check(err)

	assert.Equal(t, *fileType, internal.SZDD)
}

func TestLookupZod(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("5A 4F 4F"))
	check(err)

	assert.Equal(t, *fileType, internal.ZOD)
}

func TestLookupPbmP1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 31 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PBM_P1)
}

func TestLookupPbmP4(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 34 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PBM_P4)
}

func TestLookupPbmP2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 32 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PBM_P2)
}

func TestLookupPbmP5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 35 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PBM_P5)
}

func TestLookupPbmP3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 33 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PBM_P3)
}

func TestLookupPbmP6(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 36 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.PBM_P6)
}

func TestLookupWmf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("D7 CD C6 9A"))
	check(err)

	assert.Equal(t, *fileType, internal.WMF)
}

func TestLookupXcfGimp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("67 69 6D 70 20 78 63 66"))
	check(err)

	assert.Equal(t, *fileType, internal.XCF_GIMP)
}

func TestLookupXpm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2F 2A 20 58 50 4D 20 2A 2F"))
	check(err)

	assert.Equal(t, *fileType, internal.XPM)
}

func TestLookupAff(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("41 46 46"))
	check(err)

	assert.Equal(t, *fileType, internal.AFF)
}

func TestLookupEncaseEwfV2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("45 56 46 32"))
	check(err)

	assert.Equal(t, *fileType, internal.ENCASE_EWF_V2)
}

func TestLookupEncaseEwfV1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("45 56 46"))
	check(err)

	assert.Equal(t, *fileType, internal.ENCASE_EWF_V1)
}

func TestLookupQcow(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("51 46 49"))
	check(err)

	assert.Equal(t, *fileType, internal.QCOW)
}

func TestLookupAni(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 41 43 4F 4E"))
	check(err)

	assert.Equal(t, *fileType, internal.ANI)
}

func TestLookupCda(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 43 44 44 41"))
	check(err)

	assert.Equal(t, *fileType, internal.CDA)
}

func TestLookupQcp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 51 4C 43 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.QCP)
}

func TestLookupShockwaveDcrCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 58 ?? ?? ?? ?? 46 47 44 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.SHOCKWAVE_DCR)
}

func TestLookupShockwaveDcrCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("58 46 49 52 ?? ?? ?? ?? 4D 44 47 46"))
	check(err)

	assert.Equal(t, *fileType, internal.SHOCKWAVE_DCR)
}

func TestLookupMmDirCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 58 ?? ?? ?? ?? 4D 56 39 33"))
	check(err)

	assert.Equal(t, *fileType, internal.MM_DIR)
}

func TestLookupMmDirCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("58 46 49 52 ?? ?? ?? ?? 33 39 56 4D"))
	check(err)

	assert.Equal(t, *fileType, internal.MM_DIR)
}

func TestLookupFlv(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4C 56"))
	check(err)

	assert.Equal(t, *fileType, internal.FLV)
}

func TestLookupVdi(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3C 3C 3C 20 4F 72 61 63 6C 65 20 56 4D 20 56 69 72 74 75 61 6C 42 6F 78 20 44 69 73 6B 20 49 6D 61 67 65 20 3E 3E 3E"))
	check(err)

	assert.Equal(t, *fileType, internal.VDI)
}

func TestLookupVhd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("63 6F 6E 65 63 74 69 78"))
	check(err)

	assert.Equal(t, *fileType, internal.VHD)
}

func TestLookupVhdx(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("76 68 64 78 66 69 6C 65"))
	check(err)

	assert.Equal(t, *fileType, internal.VHDX)
}

func TestLookupIsz(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("49 73 5A 21"))
	check(err)

	assert.Equal(t, *fileType, internal.ISZ)
}

func TestLookupDaa(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("44 41 41"))
	check(err)

	assert.Equal(t, *fileType, internal.DAA)
}

func TestLookupEvt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4C 66 4C 65"))
	check(err)

	assert.Equal(t, *fileType, internal.EVT)
}

func TestLookupEvtXml(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("45 6C 66 46 69 6C 65"))
	check(err)

	assert.Equal(t, *fileType, internal.EVT_XML)
}

// func TestLookupSdb(t *testing.T) { // Offset 8
// 	fileType, err := hexsi.DetectFileType(wikiBytes("73 64 62 66"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.SDB)
// }

func TestLookupGrp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("50 4D 43 43"))
	check(err)

	assert.Equal(t, *fileType, internal.GRP)
}

func TestLookupIcm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4B 43 4D 53"))
	check(err)

	assert.Equal(t, *fileType, internal.ICM)
}

func TestLookupMsregHiv(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("72 65 67 66"))
	check(err)

	assert.Equal(t, *fileType, internal.MSREG_HIV)
}

func TestLookupMsoutlookPst(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("21 42 44 4E"))
	check(err)

	assert.Equal(t, *fileType, internal.MSOUTLOOK_PST)
}

func TestLookupDraco(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("44 52 41 43 4F"))
	check(err)

	assert.Equal(t, *fileType, internal.DRACO)
}

func TestLookupGribv1V2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("47 52 49 42"))
	check(err)

	assert.Equal(t, *fileType, internal.GRIBV1V2)
}

func TestLookupBlender(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("42 4C 45 4E 44 45 52"))
	check(err)

	assert.Equal(t, *fileType, internal.BLENDER)
}

func TestLookupJxlCase1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 00 00 0C 4A 58 4C 20 0D 0A 87 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.JXL)
}

func TestLookupJxlCase2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("FF 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.JXL)
}

func TestLookupTtf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 01 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.TTF)
}

func TestLookupOtf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4F 54 54 4F"))
	check(err)

	assert.Equal(t, *fileType, internal.OTF)
}

func TestLookupModf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 25 4D 6F 64 75 6C 65"))
	check(err)

	assert.Equal(t, *fileType, internal.MODF)
}

func TestLookupMswim(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 53 57 49 4D 00 00 00 D0 00 00 00 00"))
	check(err)

	assert.Equal(t, *fileType, internal.MSWIM)
}

func TestLookupSlob(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("21 2D 31 53 4C 4F 42 1F"))
	check(err)

	assert.Equal(t, *fileType, internal.SLOB)
}

func TestLookupSjd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("AC ED"))
	check(err)

	assert.Equal(t, *fileType, internal.SJD)
}

func TestLookupCvocf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("43 72 65 61 74 69 76 65 20 56 6F 69 63 65 20 46 69 6C 65 1A 1A 00"))
	check(err)

	assert.Equal(t, *fileType, internal.CVOCF)
}

func TestLookupAusnd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2E 73 6E 64"))
	check(err)

	assert.Equal(t, *fileType, internal.AUSND)
}

func TestLookupOglPfb(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("DB 0A CE 00"))
	check(err)

	assert.Equal(t, *fileType, internal.OGL_PFB)
}

func TestLookupHazelr(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("48 5a 4c 52 00 00 00 18"))
	check(err)

	assert.Equal(t, *fileType, internal.HAZELR)
}

func TestLookupFlp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("46 4C 68 64"))
	check(err)

	assert.Equal(t, *fileType, internal.FLP)
}

func TestLookupFlmp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("31 30 4C 46"))
	check(err)

	assert.Equal(t, *fileType, internal.FLMP)
}

func TestLookupVorencDpm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 4b 4d 43 32 31 30"))
	check(err)

	assert.Equal(t, *fileType, internal.VORENC_DPM)
}

func TestLookupMsisam(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 01 00 00 4D 53 49 53 41 4D 20 44 61 74 61 62 61 73 65"))
	check(err)

	assert.Equal(t, *fileType, internal.MSISAM)
}

func TestLookupMsaccdb(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 01 00 00 53 74 61 6E 64 61 72 64 20 41 43 45 20 44 42"))
	check(err)

	assert.Equal(t, *fileType, internal.MSACCDB)
}

func TestLookupMsmdb(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("00 01 00 00 53 74 61 6E 64 61 72 64 20 4A 65 74 20 44 42"))
	check(err)

	assert.Equal(t, *fileType, internal.MSMDB)
}

func TestLookupDrw(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("01 FF 02 04 03 02"))
	check(err)

	assert.Equal(t, *fileType, internal.DRW)
}

func TestLookupDssv2(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("02 64 73 73"))
	check(err)

	assert.Equal(t, *fileType, internal.DSSV2)
}

func TestLookupDssv3(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("03 64 73 73"))
	check(err)

	assert.Equal(t, *fileType, internal.DSSV3)
}

func TestLookupAdx(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("03 00 00 00 41 50 50 52"))
	check(err)

	assert.Equal(t, *fileType, internal.ADX)
}

func TestLookupAdobeIndd(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("06 06 ED F5 D8 1D 46 E5 BD 31 EF E7 FE 74 B7 1D"))
	check(err)

	assert.Equal(t, *fileType, internal.ADOBE_INDD)
}

func TestLookupMxf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("06 0E 2B 34 02 05 01 01 0D 01 02 01 01 02"))
	check(err)

	assert.Equal(t, *fileType, internal.MXF)
}

func TestLookupSkf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("07 53 4B 46"))
	check(err)

	assert.Equal(t, *fileType, internal.SKF)
}

func TestLookupDt2D(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("07 64 74 32 64 64 74 64"))
	check(err)

	assert.Equal(t, *fileType, internal.DT2D)
}

func TestLookupMbbtcw(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0A 16 6F 72 67 2E 62 69 74 63 6F 69 6E 2E 70 72"))
	check(err)

	assert.Equal(t, *fileType, internal.MBBTCW)
}

func TestLookupDeskmateDoc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0D 44 4F 43"))
	check(err)

	assert.Equal(t, *fileType, internal.DESKMATE_DOC)
}

func TestLookupNri(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0E 4E 65 72 6F 49 53 4F"))
	check(err)

	assert.Equal(t, *fileType, internal.NRI)
}

func TestLookupWks(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0E 57 4B 53"))
	check(err)

	assert.Equal(t, *fileType, internal.WKS)
}

func TestLookupSibMus(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0F 53 49 42 45 4C 49 55 53"))
	check(err)

	assert.Equal(t, *fileType, internal.SIB_MUS)
}

func TestLookupDsp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 20 4D 69 63 72 6F 73 6F 66 74 20 44 65 76 65 6C 6F 70 65 72 20 53 74 75 64 69 6F"))
	check(err)

	assert.Equal(t, *fileType, internal.DSP)
}

func TestLookupAmr(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 21 41 4D 52"))
	check(err)

	assert.Equal(t, *fileType, internal.AMR)
}

func TestLookupSkypeSilk(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 21 53 49 4C 4B 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.SKYPE_SILK)
}

func TestLookupRadianceHdr(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 3F 52 41 44 49 41 4E 43 45 0A"))
	check(err)

	assert.Equal(t, *fileType, internal.RADIANCE_HDR)
}

func TestLookupVbe(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 40 7E 5E"))
	check(err)

	assert.Equal(t, *fileType, internal.VBE)
}

func TestLookupCdb(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("0D F0 1D C0"))
	check(err)

	assert.Equal(t, *fileType, internal.CDB)
}

func TestLookupExtm3U(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("23 45 58 54 4D 33 55"))
	check(err)

	assert.Equal(t, *fileType, internal.EXTM3U)
}

func TestLookupM2Ar(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("6D 64 66 00"))
	check(err)

	assert.Equal(t, *fileType, internal.M2AR)
}

func TestLookupCapcomPak(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4B 50 4B 41"))
	check(err)

	assert.Equal(t, *fileType, internal.CAPCOM_PAK)
}

func TestLookupCapcomArc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("41 52 43"))
	check(err)

	assert.Equal(t, *fileType, internal.CAPCOM_ARC)
}

func TestLookupInterleafPl(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("D0 4F 50 53"))
	check(err)

	assert.Equal(t, *fileType, internal.INTERLEAF_PL)
}

// func TestLookupNifti(t *testing.T) { // Offset 344
// 	fileType, err := hexsi.DetectFileType(wikiBytes("6E 2B 31 00"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.NIFTI)
// }

// func TestLookupNiftiPair(t *testing.T) { // Offset 344
// 	fileType, err := hexsi.DetectFileType(wikiBytes("6E 69 31 00"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.NIFTI_PAIR)
// }

func TestLookupRaf64(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 41 46 36 34"))
	check(err)

	assert.Equal(t, *fileType, internal.RAF64)
}

func TestLookupVisrc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("56 49 53 33"))
	check(err)

	assert.Equal(t, *fileType, internal.VISRC)
}

func TestLookupHl7(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4D 53 48 7C 42 53 48 7C"))
	check(err)

	assert.Equal(t, *fileType, internal.HL7)
}

func TestLookupSapPwrdataV1(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("70 77 72 64 61 74 61"))
	check(err)

	assert.Equal(t, *fileType, internal.SAP_PWRDATA_V1)
}

func TestLookupArc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("1a 08"))
	check(err)

	assert.Equal(t, *fileType, internal.ARC)
}

func TestLookupPgpApubk(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2d 2d 2d 2d 2d 42 45 47 49 4e 20 50 47 50 20 50 55 42 4c 49 43 20 4b 45 49 20 42 4c 4f 43 4b 2d 2d 2d 2d 2d"))
	check(err)

	assert.Equal(t, *fileType, internal.PGP_APUBK)
}

func TestLookupCnt(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("3a 42 61 73 65 20"))
	check(err)

	assert.Equal(t, *fileType, internal.CNT)
}

func TestLookupVdrm(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 56 44 52 4d"))
	check(err)

	assert.Equal(t, *fileType, internal.VDRM)
}

func TestLookupTrid(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 59 46 46 ?? ?? ?? ?? 54 52 49 44"))
	check(err)

	assert.Equal(t, *fileType, internal.TRID)
}

func TestLookupShw4(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 73 68 77 34"))
	check(err)

	assert.Equal(t, *fileType, internal.SHW4)
}

func TestLookupShw5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 73 68 77 35"))
	check(err)

	assert.Equal(t, *fileType, internal.SHW5)
}

func TestLookupShr5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 73 68 72 35"))
	check(err)

	assert.Equal(t, *fileType, internal.SHR5)
}

func TestLookupShb5(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 73 68 62 35"))
	check(err)

	assert.Equal(t, *fileType, internal.SHB5)
}

func TestLookupRmmp(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("52 49 46 46 ?? ?? ?? ?? 52 4d 4d 50"))
	check(err)

	assert.Equal(t, *fileType, internal.RMMP)
}

func TestLookupAstmE57(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("41 53 54 4d 2d 45 35 37"))
	check(err)

	assert.Equal(t, *fileType, internal.ASTM_E57)
}

func TestLookupCrowdStrikeCf(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("aa aa aa aa"))
	check(err)

	assert.Equal(t, *fileType, internal.CROWD_STRIKE_CF)
}

func TestLookupUcas(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("8C 0A 00"))
	check(err)

	assert.Equal(t, *fileType, internal.UCAS)
}

func TestLookupUtoc(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("2D 3D 3D 2D 2D 3D 3D 2D 2D 3D 3D 2D 2D 3D 3D 2D"))
	check(err)

	assert.Equal(t, *fileType, internal.UTOC)
}

func TestLookupComodore64Bin(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("43 36 34 46 69 6C 65 00"))
	check(err)

	assert.Equal(t, *fileType, internal.COMODORE64_BIN)
}
