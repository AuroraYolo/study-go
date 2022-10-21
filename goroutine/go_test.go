package goroutine

import (
	"fmt"
	"testing"
)

func DoFmt(v int) {
	fmt.Println(v)
}
func TestGo(t *testing.T) {
	for i := 0; i < 10; i++ {
		go DoFmt(2)
		go DoFmt(1)
	}

}
