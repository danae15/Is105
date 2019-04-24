package main

import (
	"fmt"
	"github.com/danae15/ICA02/Oppgave3/b/fileutils"
)

const lang01 = ("/Users/danalexander/go/src/github.com/danae15/ICA02/Oppgave3/b/files/lang01.wl")
const lang02 = ("/Users/danalexander/go/src/github.com/danae15/ICA02/Oppgave3/b/files/lang02.wl")
const lang03 = ("/Users/danalexander/go/src/github.com/danae15/ICA02/Oppgave3/b/files/lang03.wl")

func main() {

	a := fileutils.FileToByteslice(lang01)
	b := fileutils.FileToByteslice(lang02)
	c := fileutils.FileToByteslice(lang03)

	fmt.Printf("%c \n%c \n%c", a, b, c)
}
