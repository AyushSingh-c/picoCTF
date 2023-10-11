package main

import (
	"fmt"
	"math/big"
	"rsaDecrypt"
)

func main() {
	n, _ := new(big.Int).SetString("1311097532562595991877980619849724606784164430105441327897358800116889057763413423", 10)
	c, _ := new(big.Int).SetString("861270243527190895777142537838333832920579264010533029282104230006461420086153423", 10)
	e, _ := new(big.Int).SetString("65537", 10)
	rsaInfo := rsaDecrypt.PublicInfo{n, c, e}
	plainText, _ := rsaDecrypt.DecryptRSA_FactorN(rsaInfo)
	fmt.Printf("Decrypted Plain Text: %s\n", plainText.Bytes())
}
