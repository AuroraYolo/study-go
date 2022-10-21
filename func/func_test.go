package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int { //返回值是一个匿名函数
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFunc)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarsParam(t *testing.T) { //测试可变参数
	t.Log(Sum(1, 2, 3, 4, 5))
	t.Log(Sum(1, 2, 3, 4, 5, 6))
}

func DeferClear() {
	fmt.Println("Clear resources..")
}

func TestDefer(t *testing.T) {
	defer DeferClear()
	fmt.Println("start")
	panic("error")
}
