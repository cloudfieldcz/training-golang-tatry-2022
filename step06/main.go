package main

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("start ...")
	fmt.Println(somethingSend())
}

func somethingSend() (ret bool) {
	// exception handling
	defer func() {
		if r := recover(); r != nil {
			err := r.(error)
			log.Error("log error: ", err)
			ret = false // and return false
		}
	}()

	ret = false
	if err := func1(); err != nil {
		panic(err)
	}

	// any logic there
	// ...

	if err := func2(); err != nil {
		panic(err)
	}

	// any logic there
	// ...

	if err := func3(); err != nil {
		panic(err)
	}

	// any logic there
	// ...

	ret = true
	return
}

func func1() error {
	return nil
}

func func2() error {
	return nil
}

func func3() error {
	return errors.New("Some error in my func3")
}
