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

// func TestLookupJpegRaw4(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("FF D8 FF E1 ?? ?? 45 78 69 66 00 00"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.JPEG_RAW)

// }

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

// func TestLookupIifIlbm(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 49 4C 42 4D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_ILBM)

// }

// func TestLookupIifIlbm(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 49 4C 42 4D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_ILBM)

// }

// func TestLookupIifVoice(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 38 53 56 58"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_VOICE)

// }

// func TestLookupIifAmigaCb(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 43 42 4D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_AMIGA_CB)

// }

// func TestLookupIifAniBmp(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 4E 42 4D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_ANI_BMP)

// }

// func TestLookupAniCel(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 4E 49 4D"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_ANI_CEL)

// }

// func TestLookupFaxImg(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 46 41 58 58"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_FAX_IMG)

// }

// func TestLookupIifFt(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 46 54 58 54"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_FT)

// }

// func TestLookupIifMusScoreSimple(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 53 4D 55 53"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_MUS_SCORE_SIMPLE)

// }

// func TestLookupIifMusScore(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 43 4D 55 53"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_MUS_SCORE)

// }

// func TestLookupIifYuvImage(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 59 55 56 4E"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_YUV_IMAGE)

// }

// func TestLookupAmigaFvm(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 46 41 4E 54"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_AMIGA_FVM)

// }

// func TestLookupIifAiff(t *testing.T) {
// 	fileType, err := hexsi.DetectFileType(wikiBytes("46 4F 52 4D ?? ?? ?? ?? 41 49 46 46"))
// 	check(err)

// 	assert.Equal(t, *fileType, internal.IIF_AIFF)

// }

func TestLookupLz(t *testing.T) {
	fileType, err := hexsi.DetectFileType(wikiBytes("4C 5A 49 50"))
	check(err)

	assert.Equal(t, *fileType, internal.LZ)

}
