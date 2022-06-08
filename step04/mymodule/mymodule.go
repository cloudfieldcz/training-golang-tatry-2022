package mymodule

import "fmt"

func init() {
	fmt.Println("init from mymodule!")
}

func myFunc() string {
	return "private func"
}

func MyModuleFunc() (string, string) {
	return "hello-from-mymodule", "second-string"
}
