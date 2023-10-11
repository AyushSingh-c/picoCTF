package main

import (
	"fmt"
)

func mapRune(i int) rune {
	if i >= 0 && i < 26 {
		return rune('A' + i)
	} else if i >= 26 && i < 36 {
		return rune('0' + i - 26)
	} else {
		return '_'
	}
}

func main() {
	message := []int{165, 248, 94, 346, 299, 73, 198, 221, 313, 137, 205, 87, 336, 110, 186, 69, 223, 213, 216, 216, 177, 138}
	for _, i := range message {
		fmt.Printf("%c", mapRune(i%37))
	}
}
