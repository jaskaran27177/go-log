# Logger - Simple Logging Package for Go

## Overview

Logger is a lightweight logging package written in Go, providing a simple interface for writing log messages to any `io.Writer`, such as `os.Stdout`, files, or buffers. This package is intended to be minimalistic and easy to integrate into existing projects.

## Features

- Write logs to any `io.Writer`
- Default logger writes to `os.Stdout`
- Supports structured and unstructured log messages
- Lightweight and dependency-free

---

## Usage

### Example: Logging Server Start and Errors

```go
package main

import (
	"github.com/jaskaran27177/go-log/logger"
	"github.com/jaskaran27177/go-httpmux/httpsrv"
	"net/http"
)

// tcpPort is a placeholder function for fetching the port. Replace with your actual logic.
func tcpPort() string {
	return "8080"
}

func main() {
	port := tcpPort()

	// Log the server startup message
	logger.DefaultLogger.Log("Server is running on port: ", port)

	// Start the HTTP server with the custom multiplexer
	err := http.ListenAndServe(":"+port, &httpsrv.Mux)
	if err != nil {
		logger.DefaultLogger.Log("Server got an error:", err)
	}
}
```

---

## Documentation

https://pkg.go.dev/github.com/jaskaran27177/go-log

---

## License

MIT License
