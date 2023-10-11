package main

import (
	"fmt"
	"rsaDecrypt"
)

func main() {
	rsaPublicKey := rsaDecrypt.ReadRSAPubKey("./cert")
	fmt.Println(rsaPublicKey.N)
	fmt.Println(rsaPublicKey.E)
}
