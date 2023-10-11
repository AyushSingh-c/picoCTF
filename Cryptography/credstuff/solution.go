package main

import (
	"fmt"
	"formatCTF"
)

func main() {
	cipher := "cvpbPGS{P7e1S_54I35_71Z3}"
	fmt.Println(formatCTF.DecryptRot(cipher))
}
