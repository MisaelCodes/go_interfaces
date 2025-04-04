package basic_level

import "fmt"

type ConsoleLogger struct {
    data []string
}

func (cl ConsoleLogger) Write(p []byte)(n int, err error){
    cl.data = append(cl.data, string(p))
    fmt.Printf("[LOG]: %v\n", cl.data[len(cl.data)-1])
    return len(p), nil
}
