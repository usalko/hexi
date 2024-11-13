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
	PALM_DSK_DATA      FileType = 24 //Palm Desktop Data File (Access format)
	ICON               FileType = 25 //Computer icon encoded in ICO file format
	APPLE_ICON_FORMAT  FileType = 26 //Apple Icon Image format
	THREE_GPP          FileType = 27 //3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files
	HEIC               FileType = 28 //High Efficiency Image Container (HEIC)
	Z_LZW              FileType = 29 //compressed file (often tar zip) using Lempel-Ziv-Welch algorithm
	Z_LZH              FileType = 30 //Compressed file (often tar zip) using LZH algorithm
	LZH0               FileType = 31 //Lempel Ziv Huffman archive file Method 0 (No compression)
	LZH5               FileType = 32 //Lempel Ziv Huffman archive file Method 5 (8 KiB sliding window)
	AMI_BACK           FileType = 33 //AmiBack Amiga Backup data file
	AMI_BACK_IDX       FileType = 34 //AmiBack Amiga Backup index file
	BPLIST             FileType = 35 //Binary Property List file
	BZ2                FileType = 36 //Compressed file using Bzip2 algorithm
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
	PALM_DSK_DATA: {
		Bytes:         []byte{0x00, 0x01, 0x00, 0x00},
		Offset:        0,
		NameExtension: "",
		Description:   "Palm Desktop Data File (Access format",
		Tag:           PALM_DSK_DATA,
	},
	ICON: {
		Bytes:         []byte{0x00, 0x00, 0x01, 0x00},
		Offset:        0,
		NameExtension: "ico",
		Description:   "Computer icon encoded in ICO file format",
		Tag:           ICON,
	},
	APPLE_ICON_FORMAT: {
		Bytes:         []byte{0x69, 0x63, 0x6e, 0x73},
		Offset:        0,
		NameExtension: "icns",
		Description:   "Apple Icon Image format",
		Tag:           APPLE_ICON_FORMAT,
	},
	HEIC: {
		Bytes:         []byte{0x66, 0x74, 0x79, 0x70, 0x68, 0x65, 0x69, 0x63, 0x66, 0x74, 0x79, 0x70, 0x6d},
		Offset:        4,
		NameExtension: "heic",
		Description:   "High Efficiency Image Container (HEIC)",
		Tag:           HEIC,
	},
	LZH0: {
		Bytes:         []byte{0x2D, 0x68, 0x6C, 0x30, 0x2D},
		Offset:        2,
		NameExtension: "lzh",
		Description:   "Lempel Ziv Huffman archive file Method 0 (No compression)",
		Tag:           LZH0,
	},
	LZH5: {
		Bytes:         []byte{0x2D, 0x68, 0x6C, 0x35, 0x2D},
		Offset:        2,
		NameExtension: "lzh",
		Description:   "Lempel Ziv Huffman archive file Method 5 (8 KiB sliding window)",
		Tag:           LZH5,
	},
	AMI_BACK: {
		Bytes:         []byte{0x42, 0x41, 0x43, 0x4B, 0x4D, 0x49, 0x4B, 0x45, 0x44, 0x49, 0x53, 0x4B},
		Offset:        0,
		NameExtension: "bac",
		Description:   "AmiBack Amiga Backup data file",
		Tag:           AMI_BACK,
	},
	AMI_BACK_IDX: {
		Bytes:         []byte{0x49, 0x4E, 0x44, 0x58},
		Offset:        0,
		NameExtension: "idx",
		Description:   "AmiBack Amiga Backup index file",
		Tag:           AMI_BACK_IDX,
	},
	BPLIST: {
		Bytes:         []byte{0x62, 0x70, 0x6C, 0x69, 0x73, 0x74},
		Offset:        0,
		NameExtension: "plist",
		Description:   "Binary Property List file",
		Tag:           BPLIST,
	},
	BZ2: {
		Bytes:         []byte{0x42, 0x5A, 0x68},
		Offset:        0,
		NameExtension: "bz2",
		Description:   "Compressed file using Bzip2 algorithm",
		Tag:           BZ2,
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
	THREE_GPP: {
		Bytes:         []byte{0x66, 0x74, 0x79, 0x70, 0x33, 0x67},
		Offset:        4,
		NameExtension: []string{"3gp", "3g2"},
		Description:   "3rd Generation Partnership Project 3GPP and 3GPP2 multimedia files",
		Tag:           THREE_GPP,
	},
	Z_LZW: {
		Bytes:         []byte{0x1F, 0x9D},
		Offset:        0,
		NameExtension: []string{"z", "tar.z"},
		Description:   "compressed file (often tar zip) using Lempel-Ziv-Welch algorithm",
		Tag:           Z_LZW,
	},
	Z_LZH: {
		Bytes:         []byte{0x1F, 0xA0},
		Offset:        0,
		NameExtension: []string{"z", "tar.z"},
		Description:   "Compressed file (often tar zip) using LZH algorithm",
		Tag:           Z_LZH,
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
