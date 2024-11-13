package hexsi

import (
	"fmt"

	"github.com/usalko/hexsi/internal"
)

func DetectFileType(header []byte) (*internal.FileType, error) {
	hexSignature, _ := internal.LookupSignatureByBytes1(header)
	if hexSignature != nil {
		return &hexSignature.Tag, nil
	}
	return nil, fmt.Errorf("unknown file type")
}
