package interface_test

import (
	"fmt"
	"testing"
)

type Programmer interface {
	WriteHello() string
}

type GoProgrammer struct {
}

type Pro struct {
	GoProgrammer
}

func (g *GoProgrammer) WriteHello() string {
	fmt.Println(111111)
	return "fmt.Println(\"Hello World\")"
}

func TestClient(t *testing.T) {
	var pro Pro
	pro = Pro{}
	t.Log(pro.WriteHello())
}
