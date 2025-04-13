package basic_level

import "io"

type HelloReader struct {}

func (hr HelloReader) Read(b []byte) (n int, err error){
    text := "Hello, Go"
    for i := range b{
        b[i] = byte(text[i])
    }
    return len(text), io.EOF
}
