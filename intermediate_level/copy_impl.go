package intermediate_level

import (
	"io"
	"strings"
)

type UpperCaseReader struct{
    S string
}

func (uw UpperCaseReader) Read(b []byte) (n int, err error) {
	upperCase := strings.ToUpper(uw.S)
	for i := range upperCase {
		b[i] = byte(upperCase[i])
	}
	return len(b), io.EOF
}
