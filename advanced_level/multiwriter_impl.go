package advanced_level

import (
	"fmt"
	"os"
	"time"
)

type FileLogger struct {
	path string
}

func NewFileLogger(path string) *FileLogger {
	return &FileLogger{path: path}
}

func (f FileLogger) Write(b []byte) (n int, err error) {
	d := []byte(time.Now().Format(time.DateTime))
	d = append(d, b...)
	err = os.WriteFile(f.path, d, 0644)
	return len(b), err
}

type ConsoleLogger struct{}

func (f ConsoleLogger) Write(b []byte) (n int, err error) {
	d := []byte(time.Now().Format(time.DateTime))
	d = append(d, b...)
	fmt.Println(string(d))
	return len(b), err
}
