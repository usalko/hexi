package internal

import (
	"fmt"
	"slices"
)

var lookupTable map[[8]byte][]FileType = map[[8]byte][]FileType{}

func lookupTableKey[S BytesPattern, O OffsetPattern, E NameExtensionPattern](hexSignature HexSignature[S, O, E]) [8]byte {
	switch value := any(hexSignature.Bytes); value.(type) {
	case []byte:
		return [8]byte(slices.Grow(value.([]byte), 8)[:8])
	}
	return [8]byte{}
}

func init() {
	for fileType, hexSignature := range knownSignatures1 {
		key := lookupTableKey(hexSignature)
		fileTypes, ok := lookupTable[key]
		if !ok {
			lookupTable[key] = make([]FileType, 0)
		}
		lookupTable[key] = append(fileTypes, fileType)
	}
}

func LookupSignatureByBytes1(header []byte) (*HexSignature[[]byte, uint64, string], error) {
	fileTypes, ok := lookupTable[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures1[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknow signature")
}
