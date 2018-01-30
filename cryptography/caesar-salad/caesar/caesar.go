package caesar

import "unicode"

func Decrypt(cipherText string, shifts rune) string {
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
