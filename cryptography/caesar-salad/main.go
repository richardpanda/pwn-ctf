package main

import (
	"fmt"
	"unicode"
)

func decryptCaesar(cipherText string, shifts rune) string {
	var plainText []rune

	for _, c := range cipherText {
		if unicode.IsUpper(c) {
			c = (c-65+shifts)%26 + 65
		}

		if unicode.IsLower(c) {
			c = (c-97+shifts)%26 + 97
		}

		plainText = append(plainText, c)
	}

	return string(plainText)
}

func main() {
	cipherText := "Lnl! Pynffvp EBG13 pvcure. Rnfl pelcgb, zhpu frpher. Lbh znl svaq guvf pvcure va bgure PGSf nf jryy... SYNT{pnrfne_pvcure_ebg13_e_sha}"
	fmt.Println(decryptCaesar(cipherText, 13))
}
