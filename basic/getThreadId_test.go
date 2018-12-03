package basic

import (
	"bytes"
	"strconv"
	"fmt"
	"runtime"
	"testing"
	"time"
)

//query goroutine id
func getGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

//simple task
func doTask() {
	for i := 1; i < 1000; i++ {
		var sum int
		sum = 1
		sum = sum + i
	}
	//fmt.Println(time.Since(a))

	id := getGoroutineID()
	fmt.Println("id = ", id)
}

//testing how to query goroutineId in many goroutines
func TestGetGoroutineID(t *testing.T){
	//start multi goroutine
	for i := 0;i < 10;i++{
		go doTask()
	}

	time.Sleep(5 * time.Second)
}