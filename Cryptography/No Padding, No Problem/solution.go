package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"ncCTF"
	"net"
	"strings"
)

func main() {
	// Dial the nc server
	conn, err := net.Dial("tcp", "mercury.picoctf.net:60368")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Create a buffered reader
	r := bufio.NewReader(conn)

	dict := make(map[string]string)

	ncCTF.ReadLines(r, 4)
	for i := 0; i < 3; i++ {
		temp, _ := ncCTF.ReadUntil(r, '\n')
		dict[string(temp[0])] = strings.TrimSpace(strings.Split(temp, ": ")[1])
	}
	ncCTF.ReadLines(r, 2)
	n, _ := new(big.Int).SetString(dict["n"], 10)
	c, _ := new(big.Int).SetString(dict["c"], 10)
	e, _ := new(big.Int).SetString(dict["e"], 10)

	plainText := big.NewInt(2)
	cipher := big.NewInt(2)
	cipher.Exp(plainText, e, n)
	cipher.Mul(cipher, c)

	ncCTF.SendLineAfter(conn, r, ':', cipher.String())

	temp, _ := ncCTF.ReadUntil(r, '\n')
	fmt.Println(temp)
	plainText, _ = new(big.Int).SetString(strings.TrimSpace(strings.Split(temp, ": ")[1]), 10)

	plainText.Div(plainText, big.NewInt(2))
	fmt.Printf("Decrypted Plain Text: %s\n", plainText.Bytes())
}
