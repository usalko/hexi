package internal

import (
	"fmt"
	"regexp"
	"slices"
)

var lookupTableZeroOffset map[[8]byte][]FileType = map[[8]byte][]FileType{}

var anyBytesInMiddleMasks []AnyBytesInMiddleMask = []AnyBytesInMiddleMask{
	{
		Offset: 2,
		Length: 2,
	},
	{
		Offset: 4,
		Length: 2,
	},
	{
		Offset: 4,
		Length: 4,
	},
	{
		Offset: 8,
		Length: 4,
	},
	{
		Offset: 1,
		Length: 2,
	},
	{
		Offset: 3,
		Length: 1,
	},
}

func updateLookupTableZeroOffset(key [8]byte, fileType FileType) {
	fileTypes, ok := lookupTableZeroOffset[key]
	if !ok {
		lookupTableZeroOffset[key] = make([]FileType, 0)
	}
	lookupTableZeroOffset[key] = append(fileTypes, fileType)
}

func zeroBasedKey(anyBytesInMiddle AnyBytesInMiddle) []byte {
	return slices.Concat(
		anyBytesInMiddle.Prefix,
		slices.Repeat([]byte{0x00}, int(anyBytesInMiddle.AnyBytesLength)),
		anyBytesInMiddle.Suffix,
	)
}

func lookupTableKeys[S BytesPattern, O OffsetPattern, E NameExtensionPattern](hexSignature HexSignature[S, O, E]) [][8]byte {
	result := [][8]byte{}
	switch value := any(hexSignature.Bytes); value.(type) {
	case []byte:
		result = append(result, [8]byte(slices.Grow(value.([]byte), 8)[:8]))
	case AnyBytesInMiddle:
		result = append(result, [8]byte(slices.Grow(zeroBasedKey(value.(AnyBytesInMiddle)), 8)[:8]))
	case OneOfByteSequences:
		for _, sequence := range value.(OneOfByteSequences) {
			result = append(result, [8]byte(slices.Grow(sequence, 8)[:8]))
		}
	case []AnyBytesInMiddle:
		for _, anyBytesInMiddle := range value.([]AnyBytesInMiddle) {
			result = append(result, [8]byte(slices.Grow(zeroBasedKey(anyBytesInMiddle), 8)[:8]))
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
			case OffsetMultiply:
				if any(offset).(OffsetMultiply)[0] == 0 {
					updateLookupTableZeroOffset(key, fileType)
				}
			case OffsetAny:
				updateLookupTableZeroOffset(key, fileType)
			case ZeroOrAfter:
				updateLookupTableZeroOffset(key, fileType)
			case EveryNBytes:
				if any(offset).(EveryNBytes).Initial == 0 {
					updateLookupTableZeroOffset(key, fileType)
				}
			case OffsetRange:
				if any(offset).(OffsetRange).Begin == 0 {
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
	updateLookupTables(knownSignatures01)
	updateLookupTables(knownSignatures02)
	updateLookupTables(knownSignatures03)
	updateLookupTables(knownSignatures04)
	updateLookupTables(knownSignatures05)
	updateLookupTables(knownSignatures06)
	updateLookupTables(knownSignatures07)
	updateLookupTables(knownSignatures08)
	updateLookupTables(knownSignatures09)
	updateLookupTables(knownSignatures10)
	updateLookupTables(knownSignatures11)
	updateLookupTables(knownSignatures12)
	updateLookupTables(knownSignatures13)
	updateLookupTables(knownSignatures14)
	updateLookupTables(knownSignatures15)
}

// Return 0 if all sequences are equal
func compareWithAnyBytesInMiddle(anyBytesInMiddle AnyBytesInMiddle, bytes []byte) int {
	prefixComparison := slices.Compare(anyBytesInMiddle.Prefix, bytes[:anyBytesInMiddle.AnyBytesOffset])
	if prefixComparison != 0 {
		return prefixComparison
	}
	return slices.Compare(anyBytesInMiddle.Suffix, bytes[int(anyBytesInMiddle.AnyBytesOffset+anyBytesInMiddle.AnyBytesLength):])
}

func LookupSignatureByBytes01(header []byte) (*HexSignature[[]byte, uint64, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures01[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures01[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes02(header []byte) (*HexSignature[[]byte, uint64, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures02[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures02[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes03(header []byte) (*HexSignature[OneOfByteSequences, uint64, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures03[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures03[fileType]
				for i := 0; i < len(hexSignature.Bytes); i++ {
					if ok && slices.Compare(hexSignature.Bytes[i], header) == 0 {
						return &hexSignature, nil
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes04(header []byte) (*HexSignature[OneOfByteSequences, uint64, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures04[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures04[fileType]
				for i := 0; i < len(hexSignature.Bytes); i++ {
					if ok && slices.Compare(hexSignature.Bytes[i], header) == 0 {
						return &hexSignature, nil
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes05(header []byte) (*HexSignature[AnyBytesInMiddle, uint64, []string], error) {
	for _, mask := range anyBytesInMiddleMasks {
		if len(header) < int(mask.Offset+mask.Length) {
			continue
		}
		key := zeroBasedKey(AnyBytesInMiddle{
			Prefix:         slices.Clone(header)[:mask.Offset],
			AnyBytesOffset: mask.Offset,
			AnyBytesLength: mask.Length,
			Suffix:         slices.Clone(header)[mask.Offset+mask.Length:],
		})
		fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(key, 8)[:8])]
		if ok {
			if len(fileTypes) == 1 {
				hexSignature, ok := knownSignatures05[fileTypes[0]]
				if ok {
					return &hexSignature, nil
				}
			} else {
				for _, fileType := range fileTypes {
					hexSignature, ok := knownSignatures05[fileType]
					if ok && compareWithAnyBytesInMiddle(hexSignature.Bytes, header) == 0 {
						return &hexSignature, nil
					}
				}
			}
		}
	}

	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes06(header []byte) (*HexSignature[AnyBytesInMiddle, OffsetAny, []string], error) {
	for _, mask := range anyBytesInMiddleMasks {
		if len(header) < int(mask.Offset+mask.Length) {
			continue
		}
		key := zeroBasedKey(AnyBytesInMiddle{
			Prefix:         slices.Clone(header)[:mask.Offset],
			AnyBytesOffset: mask.Offset,
			AnyBytesLength: mask.Length,
			Suffix:         slices.Clone(header)[mask.Offset+mask.Length:],
		})
		fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(key, 8)[:8])]
		if ok {
			if len(fileTypes) == 1 {
				hexSignature, ok := knownSignatures06[fileTypes[0]]
				if ok {
					return &hexSignature, nil
				}
			} else {
				for _, fileType := range fileTypes {
					hexSignature, ok := knownSignatures06[fileType]
					if ok && compareWithAnyBytesInMiddle(hexSignature.Bytes, header) == 0 {
						return &hexSignature, nil
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes07(header []byte) (*HexSignature[[]AnyBytesInMiddle, uint64, []string], error) {
	for _, mask := range anyBytesInMiddleMasks {
		if len(header) < int(mask.Offset+mask.Length) {
			continue
		}
		key := zeroBasedKey(AnyBytesInMiddle{
			Prefix:         slices.Clone(header)[:mask.Offset],
			AnyBytesOffset: mask.Offset,
			AnyBytesLength: mask.Length,
			Suffix:         slices.Clone(header)[mask.Offset+mask.Length:],
		})
		fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(key, 8)[:8])]
		if ok {
			if len(fileTypes) == 1 {
				hexSignature, ok := knownSignatures07[fileTypes[0]]
				if ok {
					return &hexSignature, nil
				}
			} else {
				for _, fileType := range fileTypes {
					hexSignature, ok := knownSignatures07[fileType]
					for _, anyBytesInMiddle := range hexSignature.Bytes {
						if ok && compareWithAnyBytesInMiddle(anyBytesInMiddle, header) == 0 {
							return &hexSignature, nil
						}
					}
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes08(header []byte) (*HexSignature[[]byte, PowerOffset, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures08[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures08[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes09(header []byte) (*HexSignature[[]byte, uint64, *regexp.Regexp], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures09[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures09[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes10(header []byte) (*HexSignature[[]byte, OffsetMultiply, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures10[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures10[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes11(header []byte) (*HexSignature[[]byte, ZeroOrAfter, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures11[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures11[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes12(header []byte) (*HexSignature[[]byte, EveryNBytes, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures12[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures12[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes13(header []byte) (*HexSignature[[]byte, OffsetFromEof, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures13[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures13[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes14(header []byte) (*HexSignature[[]byte, OffsetRange, string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures14[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures14[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}

func LookupSignatureByBytes15(header []byte) (*HexSignature[[]byte, uint64, []string], error) {
	fileTypes, ok := lookupTableZeroOffset[[8]byte(slices.Grow(slices.Clone(header), 8)[:8])]
	if ok {
		if len(fileTypes) == 1 {
			hexSignature, ok := knownSignatures15[fileTypes[0]]
			if ok {
				return &hexSignature, nil
			}
		} else {
			for _, fileType := range fileTypes {
				hexSignature, ok := knownSignatures15[fileType]
				if ok && slices.Compare(hexSignature.Bytes, header) == 0 {
					return &hexSignature, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("unknown signature")
}
