package hexsi

import (
	"fmt"

	"github.com/usalko/hexsi/internal"
)

func DetectFileType(header []byte) (*internal.FileType, error) {
	hexSignature1, _ := internal.LookupSignatureByBytes1(header)
	if hexSignature1 != nil {
		return &hexSignature1.Tag, nil
	}
	hexSignature2, _ := internal.LookupSignatureByBytes2(header)
	if hexSignature2 != nil {
		return &hexSignature2.Tag, nil
	}
	hexSignature3, _ := internal.LookupSignatureByBytes3(header)
	if hexSignature3 != nil {
		return &hexSignature3.Tag, nil
	}
	hexSignature4, _ := internal.LookupSignatureByBytes4(header)
	if hexSignature4 != nil {
		return &hexSignature4.Tag, nil
	}
	hexSignature5, _ := internal.LookupSignatureByBytes5(header)
	if hexSignature5 != nil {
		return &hexSignature5.Tag, nil
	}
	hexSignature6, _ := internal.LookupSignatureByBytes6(header)
	if hexSignature6 != nil {
		return &hexSignature6.Tag, nil
	}
	hexSignature7, _ := internal.LookupSignatureByBytes7(header)
	if hexSignature7 != nil {
		return &hexSignature7.Tag, nil
	}
	return nil, fmt.Errorf("unknown file type")
}
