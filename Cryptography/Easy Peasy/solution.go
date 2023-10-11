package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"formatCTF"
	"log"
	"ncCTF"
	"net"
	"strings"
)

func main() {
	// Dial the nc server
	conn, err := net.Dial("tcp", "mercury.picoctf.net:36981")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Create a buffered reader
	r := bufio.NewReader(conn)

	// get encrypted flag
	_ = ncCTF.ReadLines(r, 2)
	encrypted_flag, _ := ncCTF.ReadUntil(r, '\n')

	keyLength := 50000
	flagLength := len(encrypted_flag) / 2

	// get key
	input := strings.Repeat("a", keyLength-flagLength)
	ncCTF.SendLineAfter(conn, r, '?', input)
	_ = ncCTF.ReadLines(r, 1)
	key, _ := ncCTF.ReadUntil(r, '\n')

	input = strings.Repeat("a", keyLength)
	ncCTF.SendLineAfter(conn, r, '?', input)
	_ = ncCTF.ReadLines(r, 1)
	key, _ = ncCTF.ReadUntil(r, '\n')

	hexStrings := []string{encrypted_flag, hex.EncodeToString([]byte(input)), key}
	plainTextHexString, _ := formatCTF.XorHexStrings(hexStrings)
	plainTextBytes, _ := hex.DecodeString(plainTextHexString)
	fmt.Println(encrypted_flag[:10], hex.EncodeToString([]byte(input))[:10], key[:10])
	fmt.Println(string(plainTextBytes))
}
