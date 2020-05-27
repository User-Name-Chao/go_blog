package main

import (
	"fmt"
	"reflect"
)

func main() {
	d := fastrand.Uint32()
	fmt.Println("---", d, reflect.TypeOf(d).String())
}
