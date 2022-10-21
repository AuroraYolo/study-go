package main

import "fmt"

func main() {
	astr := "hello"
	aint := 1
	abool := false
	arune := 'a'
	afloat := 1.2
	astruct := struct {
	}{}
	afunc := func(op int) int {
		return op
	}

	fmt.Printf("astr 指针类型是：%T\n", &astr)
	fmt.Printf("aint 指针类型是：%T\n", &aint)
	fmt.Printf("abool 指针类型是：%T\n", &abool)
	fmt.Printf("arune 指针类型是：%T\n", &arune)
	fmt.Printf("afloat 指针类型是：%T\n", &afloat)
	fmt.Printf("astruct 指针类型是：%T\n", &astruct)
	fmt.Printf("afunc 指针类型是：%T\n", &afunc)
}
