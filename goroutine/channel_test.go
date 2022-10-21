package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func lock(c chan bool, x *int) {
	c <- true
	*x = *x + 1
	<-c
}

func TestChan(t *testing.T) {
	c := make(chan bool, 10)
	var x int
	for i := 0; i < 1; i++ {
		go lock(c, &x)
	}
	time.Sleep(time.Second)
	fmt.Println("x 的值：", x)
}

//定义只写信道类型
type Sender = chan<- int

//定义只读信道类型
type Receiver = <-chan int

func TestChannel(t *testing.T) {
	var pipline = make(chan int)

	go func() {
		var sender Sender = pipline
		fmt.Println("准备发送数据: 100")
		sender <- 100
	}()

	go func() {
		var receiver Receiver = pipline
		num := <-receiver
		fmt.Printf("接收到的数据是: %d", num)
	}()
	// 主函数sleep，使得上面两个goroutine有机会执行
	time.Sleep(time.Second)
}
