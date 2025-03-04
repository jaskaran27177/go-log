// Package logger provides a simple logging utility with a customizable writer.
//
// This package is designed to support basic logging by either writing to standard output
// or any `io.Writer` provided by the user.
package logger

import (
	"fmt"
	"io"
	"os"
)

// Logger represents a simple logger that writes logs to a specified io.Writer.
type Logger struct {
	writer io.Writer
}

// Log writes a log message to the configured writer.
//
// It accepts any number of arguments, formats them using fmt.Fprintln, and writes the
// result to the logger's output writer.
//
// Example usage:
//
//     logger.DefaultLogger.Log("Hello", "world")
//
// This will print: `Hello world` to the default writer (os.Stdout).
func (reciever Logger) Log(a ...any) {
	fmt.Fprintln(reciever.writer, a...)
}

// DefaultLogger is the default logger instance that writes to os.Stdout.
//
// Users can either use this directly or create their own Logger with a custom writer.
var DefaultLogger Logger

// init initializes the DefaultLogger to write to standard output (os.Stdout).
func init() {
	DefaultLogger.writer = os.Stdout
}
