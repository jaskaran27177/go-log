package logger

import (
	"fmt"
	"io"
	"os"
)
type Logger struct {
    writer io.Writer
}

func(reciever Logger) Log(a... any){
	fmt.Fprintln(reciever.writer, a...)
}
var DefaultLogger Logger
func init() {
	DefaultLogger.writer = os.Stdout

}