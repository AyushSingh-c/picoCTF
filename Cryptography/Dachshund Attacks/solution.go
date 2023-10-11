package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"ncCTF"
	"net"
	"rsaDecrypt"
)

func main() {
	// Dial the nc server
	conn, err := net.Dial("tcp", "mercury.picoctf.net:30761")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Create a buffered reader
	r := bufio.NewReader(conn)

	dict := make(map[string]string)

	ncCTF.ReadLines(r, 1)
	for i := 0; i < 3; i++ {
		temp, _ := ncCTF.ReadUntil(r, '\n')
		dict[string(temp[0])] = temp[3:]
	}

	n, _ := new(big.Int).SetString(dict["n"], 10)
	c, _ := new(big.Int).SetString(dict["c"], 10)
	e, _ := new(big.Int).SetString(dict["e"], 10)

	rsaInfo := rsaDecrypt.PublicInfo{N: n, C: c, E: e}
	plainText, err := rsaDecrypt.WienerAttack(rsaInfo)
	if err == nil {
		fmt.Printf("Decrypted Plain Text: %s\n", plainText.Bytes())
	} else {
		fmt.Println(err.Error())
	}
}
