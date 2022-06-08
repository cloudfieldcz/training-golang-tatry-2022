package main

import (
	"fmt"

	"cloudfield.cz/skol/step04/mymodule"
)

func main() {
	fmt.Println("start ...")
	i, j := mymodule.MyModuleFunc()
	fmt.Println(i, j)
}
