package main

import (
	"fmt"
	"reflect"
)

//var(
//	name, course, module = "DK", "Docker Deep Dive", 3.2
//)
func main() {
	name := "DK"
	//course := "Docker Deep Dive"
	module := 3.2
	ptr := &module

	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
	//fmt.Println("Memory address of *module* variable is", &module)
	fmt.Println("Memory address of *module* variable is", ptr, "and the value of *module* is", *ptr)

}
