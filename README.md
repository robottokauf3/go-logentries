# go-logentries

Logentries logging package for Go.

## Details

### Installation

    go get github.com/robottokauf3/go-logentries

### Example

    package main
    import (
        "fmt"
        "logentries"
    )
    func main() {
        logger, _ := logentries.New("1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n")
        logger.Debug("Important Debugging Message")
    }

### License

MIT

