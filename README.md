# hexi - just a database for hex file signatures

Packed and translated knowledge from [wiki](https://en.wikipedia.org/wiki/List_of_file_signatures)

Example of usage:

```go
import (
    "fmt"

    "github.com/usalko/hexi"
    "github.com/usalko/hexi/ft"
)

func main() {
    fileType, err := hexi.DetectFileType([]byte{0x52, 0x41, 0x46, 0x36, 0x34})
    if err != nil {
        panic(err)
    }

    fmt.Printf("file type is %d", fileType)
    if *fileType == ft.RAF64 {
        fmt.Print("correct detection")
    }
}
```

If you notice an error or bug please make an issue (@Thanks)