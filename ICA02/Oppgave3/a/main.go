package main

import (
	"fmt"
)

func main() {
	//
	a := "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	println("Printer ut i verb %s")
	fmt.Printf("%s\n", a)

	println("Printer ut i verb %q")
	fmt.Printf("%q\n", a)

	println("Printer ut i verb %+q")
	fmt.Printf("%+q\n", a)

	println("Printer ut i verb %c")
	fmt.Printf("%c\n", a)

	println("Printer ut i verb %c for og få svaret: 1⁄2?=? ⌘")
	fmt.Printf("%s\n", "\xc2\xbd\x3f\x3d\x3f\xe2\x8c\x98")
}
