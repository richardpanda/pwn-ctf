package main

import (
	"fmt"

	"github.com/richardpanda/pwn-ctf/cryptography/caesar-salad/caesar"
)

func main() {
	cipherText := "Lnl! Pynffvp EBG13 pvcure. Rnfl pelcgb, zhpu frpher. Lbh znl svaq guvf pvcure va bgure PGSf nf jryy... SYNT{pnrfne_pvcure_ebg13_e_sha}"
	fmt.Println(caesar.Decrypt(cipherText, 13))
}
