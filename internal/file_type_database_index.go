package internal

import (
	"fmt"
	"slices"
)

var lookupTable map[[8]byte][]FileType = map[[8]byte][]FileType{}

func lookupTableKeys[S BytesPattern, O OffsetPattern, E NameExtensionPattern](hexSignature HexSignature[S, O, E]) [][8]byte {
	result := [][8]byte{}
	switch value := any(hexSignature.Bytes); value.(type) {
	case []byte:
		result = append(result, [8]byte(slices.Grow(value.([]byte), 8)[:8]))
	case OneOfByteSequences:
		for _, sequence := range value.(OneOfByteSequences) {
			result = append(result, [8]byte(slices.Grow(sequence, 8)[:8]))
		}
	}
	return result
}

func updateLookupTable[S BytesPattern, O OffsetPattern, E NameExtensionPattern](knownSignatures map[FileType]HexSignature[S, O, E]) {
	for fileType, hexSignature := range knownSignatures {
		for _, key := range lookupTableKeys(hexSignature) {
			fileTypes, ok := lookupTable[key]
			if !ok {
				lookupTable[key] = make([]FileType, 0)
			}
			lookupTable[key] = append(fileTypes, fileType)
		}
	}
}

func init() {
	updateLookupTable(knownSignatures1)
	updateLookupTable(knownSignatures2)
	updateLookupTable(knownSignatures3)
}

func LookupSignatureByBytes1(header []byte) (*HexSignature[[]byte, uint64, string], error) {
	fileTypes, ok := lookupTable[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures1[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes2(header []byte) (*HexSignature[[]byte, uint64, []string], error) {
	fileTypes, ok := lookupTable[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures2[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes3(header []byte) (*HexSignature[OneOfByteSequences, uint64, string], error) {
	fileTypes, ok := lookupTable[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures3[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}
