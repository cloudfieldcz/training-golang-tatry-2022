package main

import "fmt"

func main() {
	fmt.Println("start ..")

	// normal struct
	x := CustomStruct{
		Name: "Test name",
	}
	fmt.Println(x.String())

	// implement abstract interface
	var a AbstractStruct = x

	fmt.Println(a.String())

}
