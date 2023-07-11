package main

import (
	"fmt"
	"io"
)

func MyGreet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
