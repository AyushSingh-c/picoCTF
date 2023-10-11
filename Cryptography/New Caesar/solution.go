package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isStringPrintable(s string) bool {
	isPrintable := true
	for _, r := range s {
		if !unicode.IsPrint(r) {
			isPrintable = false
			break
		}
	}
	return isPrintable
}

func main() {
	enc := "dcebcmebecamcmanaedbacdaanafagapdaaoabaaafdbapdpaaapadanandcafaadbdaapdpandcac"
	alpha := "abcdefghijklmnop"
	enc_shift := []string{}

	for i, _ := range alpha {
		temp := ""
		for _, c := range enc {
			index := strings.Index(alpha, string(c))
			if i <= index {
				temp += string(rune(index - i + 97))
			} else {
				temp += string(rune(index + 16 - i + 97))
			}
		}
		enc_shift = append(enc_shift, temp)
	}

	for _, b := range enc_shift {
		flag := ""
		for i := 0; i < len(b); i += 2 {
			if strings.ContainsRune(alpha, rune(b[i])) && strings.ContainsRune(alpha, rune(b[i+1])) {
				index1 := strings.Index(alpha, string(b[i]))
				index2 := strings.Index(alpha, string(b[i+1]))
				flag += string(rune(((index1 << 4) + index2)))
			}
		}
		if isStringPrintable(flag) {
			fmt.Println("Flag: ", flag)
		}
	}
}
