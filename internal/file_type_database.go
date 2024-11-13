package internal

import "regexp"

type FileType uint16

const (
	SHEBANG          FileType = 1  //Script or data to be passed to the program following the shebang (#!)
	CLARIS_WORKS     FileType = 2  //Claris Works word processing doc
	LOTUS_123_V1     FileType = 3  //Lotus 1-2-3 spreadsheet (v1) file
	LOTUS_123_V3     FileType = 4  //Lotus 1-2-3 spreadsheet (v3) file
	LOTUS_123_V4_V5  FileType = 5  //Lotus 1-2-3 spreadsheet (v4, v5) file
	LOTUS_123_V9     FileType = 6  //Lotus 1-2-3 spreadsheet (v9) file
	AMIGA_HUNK_EXE   FileType = 7  //Amiga Hunk executable file
	QUARK_EXPRESS    FileType = 8  //Quark Express document
	PASSWORD_GORILLA FileType = 9  //Password Gorilla Password Database
	LIBPCAP          FileType = 10 //Libpcap File Format
	LIBPCAP_NS       FileType = 11 //Libpcap File Format (nanosecond-resolution)
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
}

var knownSignatures2 = map[FileType]HexSignature[[]byte, uint64, []string]{
	LOTUS_123_V4_V5: {
		Bytes:         []byte{0x00, 0x00, 0x1A, 0x00, 0x02, 0x10, 0x04, 0x00, 0x00, 0x00, 0x00, 0x00},
		Offset:        0,
		NameExtension: []string{"wk4", "wk5"},
		Description:   "Lotus 1-2-3 spreadsheet (v4, v5) file",
		Tag:           LOTUS_123_V4_V5,
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
