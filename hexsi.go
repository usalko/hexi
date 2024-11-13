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
	hexSignature5, _ := internal.LookupSignatureByBytes5(header)
	if hexSignature5 != nil {
		return &hexSignature5.Tag, nil
	}
	return nil, fmt.Errorf("unknown file type")
}
