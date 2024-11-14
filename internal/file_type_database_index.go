package internal

import (
	"fmt"
	"regexp"
	"slices"
)

var lookupTableZeroOffset map[[8]byte][]FileType = map[[8]byte][]FileType{}

func updateLookupTableZeroOffset(key [8]byte, fileType FileType) {
	fileTypes, ok := lookupTableZeroOffset[key]
	if !ok {
		lookupTableZeroOffset[key] = make([]FileType, 0)
	}
	lookupTableZeroOffset[key] = append(fileTypes, fileType)
}

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

func updateLookupTables[S BytesPattern, O OffsetPattern, E NameExtensionPattern](knownSignatures map[FileType]HexSignature[S, O, E]) {
	for fileType, hexSignature := range knownSignatures {
		for _, key := range lookupTableKeys(hexSignature) {

			switch offset := hexSignature.Offset; any(offset).(type) {
			case uint64:
				if any(offset).(uint64) == 0 {
					updateLookupTableZeroOffset(key, fileType)
				}
			case PowerOffset:
				if any(offset).(PowerOffset).Initial == 0 {
					updateLookupTableZeroOffset(key, fileType)
				}
			}
		}
	}
}

func init() {
	updateLookupTables(knownSignatures1)
	updateLookupTables(knownSignatures2)
	updateLookupTables(knownSignatures3)
	updateLookupTables(knownSignatures4)
	updateLookupTables(knownSignatures5)
	updateLookupTables(knownSignatures6)
	updateLookupTables(knownSignatures7)
	updateLookupTables(knownSignatures8)
	updateLookupTables(knownSignatures9)
}

func LookupSignatureByBytes1(header []byte) (*HexSignature[[]byte, uint64, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures1[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes2(header []byte) (*HexSignature[[]byte, uint64, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures2[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes3(header []byte) (*HexSignature[OneOfByteSequences, uint64, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures3[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes4(header []byte) (*HexSignature[OneOfByteSequences, uint64, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures4[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes5(header []byte) (*HexSignature[AnyBytesInMiddle, uint64, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures5[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes6(header []byte) (*HexSignature[AnyBytesInMiddle, OffsetAny, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures6[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes7(header []byte) (*HexSignature[[]byte, PowerOffset, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures7[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes8(header []byte) (*HexSignature[[]byte, uint64, *regexp.Regexp], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures8[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes9(header []byte) (*HexSignature[[]byte, uint64, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok && len(fileTypes) == 1 {
		hexSignature, ok := knownSignatures9[fileTypes[0]]
		if ok {
			return &hexSignature, nil
		}
	}
	return nil, fmt.Errorf("unknown signature")
}
