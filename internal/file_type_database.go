package internal

import "regexp"

type FileType uint16

const (
	SHEBANG FileType = 1
)

type AnyBytesInMiddle struct {
	Prefix         []byte
	AnyBytesOffset uint8
	AnyBytesLength uint8
	Suffix         []byte
}

type BytesPattern interface {
	[]byte | AnyBytesInMiddle
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
}
