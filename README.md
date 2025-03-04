# logger

A simple Go logging package with customizable output writer support.

## Installation

To install this package, run:

```bash
go get github.com/jaskaran27177/go-log
```

## Usage

```go
package main

import (
	"github.com/jaskaran27177/go-log/logger"
	"bytes"
)

func main() {
	// Log using the default logger (writes to stdout)
	logger.DefaultLogger.Log("This is a log message")

	// Create a custom logger
	var buf bytes.Buffer
	customLogger := logger.Logger{Writer: &buf}
	customLogger.Log("Log to buffer")
}
```

## Package Documentation

You can view the full documentation on **pkg.go.dev**:

🔗 [https://pkg.go.dev/github.com/jaskaran27177/go-log](https://pkg.go.dev/github.com/jaskaran27177/go-log)

## Features

- Simple logging interface
- Supports any `io.Writer`
- Default logger writes to `os.Stdout`

## License

MIT License
