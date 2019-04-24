package treasure

import (
	"bytes"
	"fmt"
)

// PrintTreasureUTF8 tar en "string literal" som INN-data
// og skriver ut en korrekt text på med UTF-8 koding
// Koden er for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) {
	treasureByte := []byte(treasureString)

	//Replaces UFT16 Ø with UFT8 Ø
	treasureByte = bytes.Replace(treasureByte, []byte("\xf8"), []byte("\xc3s\xb8"), 1)
	//Replaces UFT16 Æ with UFT8 Æ
	treasureByte = bytes.Replace(treasureByte, []byte("\xe6"), []byte("\xc3\xa6"), 1)
	//Replaces UFT16 Å with UFT8 Å
	treasureByte = bytes.Replace(treasureByte, []byte("\xe5"), []byte("\xc3\xa5"), 1)
	fmt.Printf("%s", treasureByte)
}
