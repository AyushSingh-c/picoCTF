package main

import (
	"fmt"
	"math/big"
)

func mapRune(i int) rune {
	if i >= 1 && i <= 26 {
		return rune('A' + i - 1)
	} else if i > 26 && i <= 36 {
		return rune('0' + i - 27)
	} else {
		return '_'
	}
}

func main() {
	message := []int{268, 413, 438, 313, 426, 337, 272, 188, 392, 338, 77, 332, 139, 113, 92, 239, 247, 120, 419, 72, 295, 190, 131}

	for _, i := range message {
		// fmt.Println(new(big.Int).ModInverse(big.NewInt(int64(i%41)), big.NewInt(41)).Int64())
		fmt.Printf("%c", mapRune(int(new(big.Int).ModInverse(big.NewInt(int64(i%41)), big.NewInt(41)).Int64())))
	}
}
