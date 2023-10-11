package main

import (
	"bufio"
	"cryptoES"
	"encoding/hex"
	"fmt"
	"log"
	"ncCTF"
	"net"
	"strconv"
)

func main() {

	// Dial the nc server
	conn, err := net.Dial("tcp", "mercury.picoctf.net:5958")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Create a buffered reader
	r := bufio.NewReader(conn)

	// Sample encrypted and decrypted pair
	ncCTF.ReadLines(r, 1)
	flagCLI, _ := ncCTF.ReadUntil(r, '\n')
	flag_enc, _ := hex.DecodeString(flagCLI)
	plaintext := "abcd    "
	ncCTF.SendLineAfter(conn, r, '?', "61626364")
	encryptionCLI, _ := ncCTF.ReadUntil(r, '\n')
	encryption, _ := hex.DecodeString(encryptionCLI)

	fmt.Println("Flag: ", flagCLI)
	fmt.Println("plaintext: ", plaintext)
	fmt.Println("encryption: ", encryptionCLI)

	listKey := []string{}

	// Loop from 0 to 999999
	for i := 0; i <= 999999; i++ {
		s := strconv.Itoa(i)
		s = fmt.Sprintf("%06s", s)
		key := []byte(s + "  ")
		listKey = append(listKey, string(key))
	}
	cryptoES.Solve2DES([]byte(plaintext), encryption, flag_enc, listKey)
}
