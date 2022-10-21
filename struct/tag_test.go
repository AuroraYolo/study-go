package _struct

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name string `json:"name" label:"hello"`
	Age  int    `json:"age"`
	Addr string `json:"addr"`
}

func TestTag(t *testing.T) {
	p1 := Person{
		Name: "Jack",
		Age:  22,
	}

	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}

	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)

	field, _ := reflect.TypeOf(p1).FieldByName("Name")
	fmt.Println(field.Tag.Get("label"))

}
