package hexsi

import (
	"fmt"

	"github.com/usalko/hexsi/ft"
)

func DetectFileType(header []byte) (*ft.FileType, error) {
	hexSignature01, _ := ft.LookupSignatureByBytes01(header)
	if hexSignature01 != nil {
		return &hexSignature01.Tag, nil
	}
	hexSignature02, _ := ft.LookupSignatureByBytes02(header)
	if hexSignature02 != nil {
		return &hexSignature02.Tag, nil
	}
	hexSignature03, _ := ft.LookupSignatureByBytes03(header)
	if hexSignature03 != nil {
		return &hexSignature03.Tag, nil
	}
	hexSignature04, _ := ft.LookupSignatureByBytes04(header)
	if hexSignature04 != nil {
		return &hexSignature04.Tag, nil
	}
	hexSignature05, _ := ft.LookupSignatureByBytes05(header)
	if hexSignature05 != nil {
		return &hexSignature05.Tag, nil
	}
	hexSignature06, _ := ft.LookupSignatureByBytes06(header)
	if hexSignature06 != nil {
		return &hexSignature06.Tag, nil
	}
	hexSignature07, _ := ft.LookupSignatureByBytes07(header)
	if hexSignature07 != nil {
		return &hexSignature07.Tag, nil
	}
	hexSignature08, _ := ft.LookupSignatureByBytes08(header)
	if hexSignature08 != nil {
		return &hexSignature08.Tag, nil
	}
	hexSignature09, _ := ft.LookupSignatureByBytes09(header)
	if hexSignature09 != nil {
		return &hexSignature09.Tag, nil
	}
	hexSignature10, _ := ft.LookupSignatureByBytes10(header)
	if hexSignature10 != nil {
		return &hexSignature10.Tag, nil
	}
	hexSignature11, _ := ft.LookupSignatureByBytes11(header)
	if hexSignature11 != nil {
		return &hexSignature11.Tag, nil
	}
	hexSignature12, _ := ft.LookupSignatureByBytes12(header)
	if hexSignature12 != nil {
		return &hexSignature12.Tag, nil
	}
	hexSignature13, _ := ft.LookupSignatureByBytes13(header)
	if hexSignature13 != nil {
		return &hexSignature13.Tag, nil
	}
	hexSignature14, _ := ft.LookupSignatureByBytes14(header)
	if hexSignature14 != nil {
		return &hexSignature14.Tag, nil
	}
	hexSignature15, _ := ft.LookupSignatureByBytes15(header)
	if hexSignature15 != nil {
		return &hexSignature15.Tag, nil
	}
	return nil, fmt.Errorf("unknown file type")
}
