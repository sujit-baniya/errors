package main

import (
	stdErrors "errors"
	"fmt"
	"githubc.com/sujit-baniya/errors"
)

func main() {
	err := A()
	fmt.Println(errors.ErrorStack(err))
	// fmt.Println(errors.RedactLogEnabled)
}

func A() error {
	err := B()
	if err != nil {
		return errors.Wrap(err, "Found error on B")
	}
	return errors.New("Error on A")
}

func B() error {
	err := C()
	if err != nil {
		return errors.Wrap(err, "Found error on C")
	}
	return stdErrors.New("Error on B")
}

func C() error {
	return stdErrors.New("Error on C")
}
