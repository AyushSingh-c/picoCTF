package main

import (
	"fmt"
	"strings"
)

// decodeRailFenceCipher takes a cipher text and a number of rails
// and returns the plain text using the rail-fence cipher algorithm
func decodeRailFenceCipher(cipher string, rails int) string {
	// create a matrix of runes to store the cipher and plain text
	matrix := make([][]rune, rails)
	for i := range matrix {
		matrix[i] = make([]rune, len(cipher))
	}

	// mark the spots where the cipher text should be placed
	dirDown := false // direction of writing
	row, col := 0, 0 // current row and column
	for _, ch := range cipher {
		// change direction when we reach the top or bottom rail
		if row == 0 || row == rails-1 {
			dirDown = !dirDown
		}
		// place a dot at the current spot
		matrix[row][col] = '.'
		// move to the next column
		col++
		// move to the next row depending on the direction
		if dirDown {
			row++
		} else {
			row--
		}
		fmt.Print(ch)
	}
	fmt.Print("\n")
	// fill the cipher text in the marked spots
	index := 0 // index of the cipher text
	for i := 0; i < rails; i++ {
		for j := 0; j < len(cipher); j++ {
			if matrix[i][j] == '.' && index < len(cipher) {
				matrix[i][j] = rune(cipher[index])
				index++
			}
		}
	}
	dirDown = false // direction of writing
	// read the plain text in a zig-zag manner
	var plain strings.Builder // plain text builder
	row, col = 0, 0           // current row and column
	for col < len(cipher) {
		if row == 0 || row == rails-1 {
			dirDown = !dirDown
		}
		if matrix[row][col] != 0 {
			fmt.Fprintf(&plain, "%c", matrix[row][col])
			col++
		}
		if dirDown {
			row++
		} else {
			row--
		}
	}

	return plain.String()
}

func main() {
	cipher := "Ta _7N6D49hlg:W3D_H3C31N__A97ef sHR053F38N43D7B i33___N6"
	rails := 4
	fmt.Println("Cipher text:", cipher)
	fmt.Println("Number of rails:", rails)
	fmt.Println("Plain text:", decodeRailFenceCipher(cipher, rails))
}
