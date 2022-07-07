package main

import (
	"fmt"

	"github.com/tinystack/errors"
)

func main() {
	err1 := exampleError()
	fmt.Println(err1.Error())
	// main.exampleError(main.go:28): example err

	err2 := exampleErrorf()
	fmt.Println(err2.Error())
	// main.exampleErrorf(main.go:32): example err: error message

	err3 := exampleWrap()
	fmt.Println(err3.Error())
	// main.exampleWrap(main.go:37): wrap error message | Caused: main.exampleWrap(main.go:36): simple err

	err4 := exampleWrapf()
	fmt.Println(err4.Error())
	// main.exampleWrapf(main.go:42): wrap: error message | Caused: main.exampleWrapf(main.go:41): simple err
}

func exampleError() error {
	return errors.New("example err")
}

func exampleErrorf() error {
	return errors.Newf("example err: %s", "error message")
}

func exampleWrap() error {
	err := errors.New("simple err")
	return errors.Wrap(err, "wrap error message")
}

func exampleWrapf() error {
	err := errors.New("simple err")
	return errors.Wrapf(err, "wrap: %s", "error message")
}
