package main

import (
	"fmt"

	"github.com/kran891/go-lang-api/common/errors"
)

func main() {
	k := errors.New("Unknown Error")
	z := errors.Wrap(k, "Hello")
	m := errors.UnWrap(z)
	fmt.Print(m)
}
