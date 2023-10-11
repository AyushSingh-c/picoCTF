package main

import (
	"fmt"
	"strings"
)

// rot13 applies the ROT13 cipher to a rune
func rot13(r rune) rune {
	// Check if the rune is a letter and its case
	isLetter, isUpper := letterCase(r)
	if !isLetter {
		return r // Not a letter, do nothing
	}
	// Rotate the rune by 13 places
	r += 13
	// If the rune goes outside the alphabet, wrap it around
	if isUpper && r > 'Z' || !isUpper && r > 'z' {
		r -= 26
	}
	return r
}

// letterCase checks if a rune is a letter and returns its case
func letterCase(r rune) (bool, bool) {
	isUpper := r >= 'A' && r <= 'Z'
	isLower := r >= 'a' && r <= 'z'
	return isUpper || isLower, isUpper
}

// encrypt encrypts a string using ROT13
func encrypt(s string) string {
	return strings.Map(rot13, s)
}

// decrypt decrypts a string using ROT13
func decrypt(s string) string {
	return strings.Map(rot13, s) // Same as encrypt, since ROT13 is reversible
}

func main() {
	// Test the program with some examples
	fmt.Println(decrypt("cvpbPGS{arkg_gvzr_V'yy_gel_2_ebhaqf_bs_ebg13_hyLicInt}")) // Uryyb, jbeyq!
}
