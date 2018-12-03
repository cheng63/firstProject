package log

import (
	"fmt"
	"bytes"
	"testing"
)



func TestCli(t *testing.T) {
	b := &bytes.Buffer{}
	//fmt.Fprintf(b, "\x1b[%dm%s\x1b[0m[%s] %s ", 35, "INFO ", "01-02|15:04:05", "erongdu")

	fmt.Fprintf(b, "\x1b[%dm lvl = %s\x1b[32m time = %s msg = %s", 35, "INFO ","01-02|15:04:05 ","erongdu")
	//fmt.Fprintf(b, "\x1b[0m time = %s msg = %s", "01-02|15:04:05 ","erongdu")

	//fmt.Fprintf(b, "\x1b[%dm lvl = %s", 32, "INFO ")
	//fmt.Fprintf(b, "\x1b[0m time = %s msg = %s", "01-02|15:04:05 ","erongdu")
	fmt.Println(b)
}