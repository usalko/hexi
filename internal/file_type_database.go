package internal

import "regexp"

type FileType uint16

const (
	SHEBANG            FileType = 1  //Script or data to be passed to the program following the shebang (#!)
	CLARIS_WORKS       FileType = 2  //Claris Works word processing doc
	LOTUS_123_V1       FileType = 3  //Lotus 1-2-3 spreadsheet (v1) file
	LOTUS_123_V3       FileType = 4  //Lotus 1-2-3 spreadsheet (v3) file
	LOTUS_123_V4_V5    FileType = 5  //Lotus 1-2-3 spreadsheet (v4, v5) file
	LOTUS_123_V9       FileType = 6  //Lotus 1-2-3 spreadsheet (v9) file
	AMIGA_HUNK_EXE     FileType = 7  //Amiga Hunk executable file
	QUARK_EXPRESS      FileType = 8  //Quark Express document
	PASSWORD_GORILLA   FileType = 9  //Password Gorilla Password Database
	LIBPCAP            FileType = 10 //Libpcap File Format
	LIBPCAP_NS         FileType = 11 //Libpcap File Format (nanosecond-resolution)
	PCAPNPG            FileType = 12 //PCAP Next Generation Dump File Format
	RPM                FileType = 13 //RedHat Package Manager (RPM) package
	SQLITE3            FileType = 14 //SQLite Database
	AMAZON_KINDLE_UP   FileType = 15 //Amazon Kindle Update Package
	DOOM_WAD           FileType = 16 //internal WAD (main resource file of Doom)
	ZERO               FileType = 17 //IBM Storyboard bitmap file, Windows Program Information File, Mac Stuffit Self-Extracting Archive, IRIS OCR data file
	PALM_PILOT         FileType = 18 //PalmPilot Database/Document File
	PALM_DSK_CALENDAR  FileType = 19 //Palm Desktop Calendar Archive
	PALM_DSK_TODO      FileType = 20 //Palm Desktop To Do Archive
	PALM_DSK_CALENDAR2 FileType = 21 //Palm Desktop Calendar Archive
	TELEGRAM_DSK       FileType = 22 //Telegram Desktop File
	TELEGRAM_DSK_ENC   FileType = 23 //Telegram Desktop Encrypted File
)

type AnyBytesInMiddle struct {
	Prefix         []byte
	AnyBytesOffset uint8
	AnyBytesLength uint8
	Suffix         []byte
}

type OneOfByteSequences [][]byte

type BytesPattern interface {
	[]byte | AnyBytesInMiddle | OneOfByteSequences
}

type OffsetFromEof uint64
type OffsetAny bool
type OffsetMultiply []uint64

type OffsetPattern interface {
	uint64 | OffsetFromEof | OffsetAny | OffsetMultiply
}

type NameExtensionPattern interface {
	string | []string | regexp.Regexp
}

type HexSignature[S BytesPattern, O OffsetPattern, E NameExtensionPattern] struct {
	Bytes         S
	Offset        O
	NameExtension E
	Description   string
	Tag           FileType
}

var knownSignatures1 = map[FileType]HexSignature[[]byte, uint64, string]{
	SHEBANG: {
		Bytes:         []byte{0x23, 0x21},
		Offset:        0,
		NameExtension: "",
		Description:   "Script or data to be passed to the program following the shebang (#!)",
		Tag:           SHEBANG,
	},
	CLARIS_WORKS: {
		Bytes:         []byte{0x02, 0x00, 0x5a, 0x57, 0x52, 0x54, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "cwk",
		Description:   "Claris Works word processing doc",
		Tag:           CLARIS_WORKS,
	},
	LOTUS_123_V1: {
		Bytes:         []byte{0x00, 0x00, 0x02, 0x00, 0x06, 0x04, 0x06, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "wk1",
		Description:   "Lotus 1-2-3 spreadsheet (v1) file",
		Tag:           LOTUS_123_V1,
	},
	LOTUS_123_V3: {
		Bytes:         []byte{0x00, 0x00, 0x1A, 0x00, 0x00, 0x10, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: "wk3",
		Description:   "Lotus 1-2-3 spreadsheet (v3) file",
		Tag:           LOTUS_123_V3,
	},
	LOTUS_123_V9: {
		Bytes:         []byte{0x00, 0x00, 0x1A, 0x00, 0x05, 0x10, 0x04},
		Offset:        0,
		NameExtension: "123",
		Description:   "Lotus 1-2-3 spreadsheet (v9) file",
		Tag:           LOTUS_123_V9,
	},
	AMIGA_HUNK_EXE: {
		Bytes:         []byte{0x00, 0x00, 0x03, 0xF3},
		Offset:        0,
		NameExtension: "",
		Description:   "Amiga Hunk executable file",
		Tag:           AMIGA_HUNK_EXE,
	},
	PASSWORD_GORILLA: {
		Bytes:         []byte{0x50, 0x57, 0x53, 0x33},
		Offset:        0,
		NameExtension: "psafe3",
		Description:   "Password Gorilla Password Database",
		Tag:           PASSWORD_GORILLA,
	},
	PCAPNPG: {
		Bytes:         []byte{0x0A, 0x0D, 0x0D, 0x0A},
		Offset:        0,
		NameExtension: "pcapnpg",
		Description:   "PCAP Next Generation Dump File Format",
		Tag:           PCAPNPG,
	},
	RPM: {
		Bytes:         []byte{0xED, 0xAB, 0xEE, 0xDB},
		Offset:        0,
		NameExtension: "rpm",
		Description:   "RedHat Package Manager (RPM) package",
		Tag:           RPM,
	},
	AMAZON_KINDLE_UP: {
		Bytes:         []byte{0x53, 0x50, 0x30, 0x31},
		Offset:        0,
		NameExtension: "bin",
		Description:   "Amazon Kindle Update Package",
		Tag:           AMAZON_KINDLE_UP,
	},
	DOOM_WAD: {
		Bytes:         []byte{0x49, 0x57, 0x41, 0x44},
		Offset:        0,
		NameExtension: "wad",
		Description:   "internal WAD (main resource file of Doom)",
		Tag:           DOOM_WAD,
	},
	PALM_PILOT: {
		Bytes:         []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        11,
		NameExtension: "PDB",
		Description:   "PalmPilot Database/Document File",
		Tag:           PALM_PILOT,
	},
	PALM_DSK_CALENDAR: {
		Bytes:         []byte{0xBE, 0xBA, 0xFE, 0xCA},
		Offset:        0,
		NameExtension: "DBA",
		Description:   "Palm Desktop Calendar Archive",
		Tag:           PALM_DSK_CALENDAR,
	},
	PALM_DSK_TODO: {
		Bytes:         []byte{0x00, 0x01, 0x42, 0x44},
		Offset:        0,
		NameExtension: "DBA",
		Description:   "Palm Desktop To Do Archive",
		Tag:           PALM_DSK_TODO,
	},
	PALM_DSK_CALENDAR2: {
		Bytes:         []byte{0x00, 0x01, 0x44, 0x54},
		Offset:        0,
		NameExtension: "DBA",
		Description:   "Palm Desktop Calendar Archive",
		Tag:           PALM_DSK_CALENDAR2,
	},
	TELEGRAM_DSK: {
		Bytes:         []byte{0x54, 0x44, 0x46, 0x24},
		Offset:        0,
		NameExtension: "TDF$",
		Description:   "Telegram Desktop File",
		Tag:           TELEGRAM_DSK,
	},
	TELEGRAM_DSK_ENC: {
		Bytes:         []byte{0x54, 0x44, 0x45, 0x46},
		Offset:        0,
		NameExtension: "TDEF",
		Description:   "Telegram Desktop Encrypted File",
		Tag:           TELEGRAM_DSK_ENC,
	},
}

var knownSignatures2 = map[FileType]HexSignature[[]byte, uint64, []string]{
	LOTUS_123_V4_V5: {
		Bytes:         []byte{0x00, 0x00, 0x1A, 0x00, 0x02, 0x10, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: []string{"wk4", "wk5"},
		Description:   "Lotus 1-2-3 spreadsheet (v4, v5) file",
		Tag:           LOTUS_123_V4_V5,
	},
	SQLITE3: {
		Bytes:         []byte{0x53, 0x51, 0x4C, 0x69, 0x74, 0x65, 0x20, 0x66, 0x6F, 0x72, 0x6D, 0x61, 0x74, 0x20, 0x33, 0x00},
		Offset:        0,
		NameExtension: []string{"sqlitedb", "sqlite", "db"},
		Description:   "SQLite Database",
		Tag:           SQLITE3,
	},
}

var knownSignatures3 = map[FileType]HexSignature[OneOfByteSequences, uint64, string]{
	QUARK_EXPRESS: {
		Bytes: [][]byte{
			{0x00, 0x00, 0x49, 0x49, 0x58, 0x50, 0x52},
			{0x00, 0x00, 0x4D, 0x4D, 0x58, 0x50, 0x52},
		},
		Offset:        0,
		NameExtension: "qxd",
		Description:   "Quark Express document",
		Tag:           QUARK_EXPRESS,
	},
	LIBPCAP: {
		Bytes: [][]byte{
			{0xD4, 0xC3, 0xB2, 0xA1},
			{0xA1, 0xB2, 0xC3, 0xD4},
		},
		Offset:        0,
		NameExtension: "pcap",
		Description:   "Libpcap File Format",
		Tag:           LIBPCAP,
	},
	LIBPCAP_NS: {
		Bytes: [][]byte{
			{0x4D, 0x3C, 0xB2, 0xA1},
			{0xA1, 0xB2, 0x3C, 0x4D},
		},
		Offset:        0,
		NameExtension: "pcap",
		Description:   "Libpcap File Format (nanosecond-resolution)",
		Tag:           LIBPCAP_NS,
	},
}

var knownSignatures5 = map[FileType]HexSignature[[]byte, uint64, []string]{
	ZERO: {
		Bytes:         []byte{0x00},
		Offset:        0,
		NameExtension: []string{"PIC", "PIF", "SEA", "YTR"},
		Description:   "IBM Storyboard bitmap file, Windows Program Information File, Mac Stuffit Self-Extracting Archive, IRIS OCR data file",
		Tag:           ZERO,
	},
}
