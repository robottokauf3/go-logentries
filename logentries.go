// Go support for logging to Logentries, http://logentries.com
// Version 0.1.0
//
// Copyright 2014 Robert Kaufmann III

package logentries

import (
	"errors"
	"fmt"
	"log"
	"net"
)

// Log levels
type Level int

const (
	Debug Level = iota
	Info
	Warn
	Error
	Panic
)

// Logentires endpoint
const (
	endpoint string = "data.logentries.com:80"
)

// Default tags for log levels
var levelTag = map[Level]string{
	Debug: "DEBUG",
	Info:  "INFO",
	Warn:  "WARNING",
	Error: "ERROR",
	Panic: "PANIC",
}

// Log Client
type Logger struct {
	localLogger *log.Logger
	connection  net.Conn
	token       string
	verbosity   Level
}

// New creates a new logging client with supplied token
func New(token string) (*Logger, error) {
	// Setup the connection
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		return nil, err
	}

	return &Logger{
		token:       token,
		verbosity:   0,
		localLogger: nil,
		connection:  conn,
	}, nil
}

// SetVerbosity sets the minimal log level which will be recorded
func (l *Logger) SetVerbosity(level Level) error {
	if _, ok := levelTag[level]; ok {
		l.verbosity = level
		return nil
	}
	return errors.New("unknown log level")
}

// SetLocalLogger sets a local log instance
// the log message will be passed to.
func (l *Logger) SetLocalLogger(local *log.Logger) {
	l.localLogger = local
}

// Debug is a helper function
func (l *Logger) Debug(message string) error {
	return l.Log(Debug, message)
}

// Info is a helper function
func (l *Logger) Info(message string) error {
	return l.Log(Info, message)
}

// Warn is a helper function
func (l *Logger) Warn(message string) error {
	return l.Log(Warn, message)
}

// Error is a helper function
func (l *Logger) Error(message string) error {
	return l.Log(Error, message)
}

// Panic is a helper function
func (l *Logger) Panic(message string) error {
	return l.Log(Panic, message)
}

// Log sends a formatted log message logentries
// if the log is higher than the current minimium level.
// If a local logger is set then Log also passes the message to it.
func (l *Logger) Log(level Level, message string) error {
	if level >= l.verbosity {
		if l.localLogger != nil {
			log.Print(message)
		}
		remoteMessage := l.format(level, message)
		return l.send(remoteMessage)
	}

	return nil
}

// format prepares the log for transmitted by
// prepending the UUID token and appending a new line.
func (l *Logger) format(level Level, message string) string {
	return l.token + " " + levelTag[level] + " " + message + "\n"
}

// SendRaw sends an unformatted log
func (l *Logger) SendRaw(message string) error {
	log := l.token + " " + message + "\n"
	return l.send(log)
}

// send transmits message to LogEntries
func (l *Logger) send(log string) error {
	fmt.Println(log)
	_, err := l.connection.Write([]byte(log))
	return err
}
