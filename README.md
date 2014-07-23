# go-logentries

Logentries logging package for Go. The token-based input is used with SSL available when required.

## Details

### Installation

    go get github.com/robottokauf3/go-logentries

### Example

    package main

    import (
        "fmt"
        "github.com/robottokauf3/go-logentries"
    )

    func main() {
        // Using standard TCP connection
        logger, _ := logentries.New("1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n", false)
        logger.Debug("Important Debugging Message")

        // Using SSL connection
        sslLogger, _ := logentries.New("1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n", true)
        sslLogger.Debug("Secret Debugging Message")

        // Setting the minimum log level
        warningLogger, _ := logentries.New("1a2b3c4d-5e6f-7g8h-9i0j-1k2l3m4n", true)
        warningLogger.SetVerbosity(logentries.Warn);
        warningLogger.Debug("I won't be logged")
        warningLogger.Error("But I will")
    }

### License

MIT

