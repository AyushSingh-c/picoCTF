package main

import (
	"bufio"
	"fmt"
	"log"
	"ncCTF"
	"net"
	"strings"
)

func getEnc(s string, lis []string) string {
	temp := s
	for _, l := range lis {
		temp = strings.ReplaceAll(temp, l, "")
	}
	return temp
}

func main() {

	// Dial the nc server
	conn, err := net.Dial("tcp", "mercury.picoctf.net:4484")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Create a buffered reader
	r := bufio.NewReader(conn)

	// Sample encrypted and decrypted pair
	possible := "picoCTF{bad_1d3a5_5700361}" + "qwertyuiopasdfghjklzxcvbnm_{}QWERTYUIOPASDFGHJKLZXCVBNM0123456789"
	flag := ""
	listEnc := []string{}
	flagEnc, _ := ncCTF.ReadUntil(r, '\n')

	fmt.Println(flagEnc)
	ncCTF.ReadLines(r, 2)

	for {
		for _, c := range possible {
			ncCTF.SendLineAfter(conn, r, ':', flag+string(c))
			ncCTF.ReadUntil(r, ':')
			temp, _ := ncCTF.ReadUntil(r, '\n')
			temp = getEnc(temp, listEnc)
			// fmt.Printf("char: %s\nenc: %s\n", flag+string(c), temp)
			if strings.Contains(flagEnc, temp) {
				flag += string(c)
				fmt.Printf("Flag until now: %s\nNew char: %c\n", flag, c)
				listEnc = append(listEnc, temp)
				break
			}
		}
	}

}
