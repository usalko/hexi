# hexsi - just a database for hex file signatures

Packed and translated knowledge from [wiki](https://en.wikipedia.org/wiki/List_of_file_signatures)

Example of usage:

```lang:go
import (
    "fmt"

    "github.com/usalko/hexsi"
    "github.com/usalko/hexsi/internal"
)

func main() {
    fileType, err := hexsi.DetectFileType([]byte{0x52, 0x41, 0x46, 0x36, 0x34})
    if err != nil {
        panic(err)
    }

    fmt.Printf("file type is %d", fileType)
    if *fileType == internal.RAF64 {
        fmt.Print("correct detection")
    }
}
```

If you notice an error or bug please make an issue (@Thanks)