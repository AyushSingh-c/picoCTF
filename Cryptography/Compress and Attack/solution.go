package main

import (
	"bufio"
	"fmt"
	"log"
	"ncCTF"
	"net"
	"strconv"
)

func main() {

	// Dial the nc server
	conn, err := net.Dial("tcp", "mercury.picoctf.net:2431")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	// Create a buffered reader
	r := bufio.NewReader(conn)

	// Sample encrypted and decrypted pair
	flag := "picoCTF{sheriff_you_solved_the_crime}"
	possible := "qwertyuiopasdfghjklzxcvbnm_{}QWERTYUIOPASDFGHJKLZXCVBNM"

	lengthComp := func(text string, conn net.Conn, r *bufio.Reader) int {
		ncCTF.SendLineAfter(conn, r, ':', text+text)
		ncCTF.ReadLines(r, 2)
		temp, _ := ncCTF.ReadUntil(r, '\n')
		value, _ := strconv.Atoi(temp)
		return value
	}

	for {
		final := '?'
		value := lengthComp(flag+string(final), conn, r)
		bre := 0
		for _, c := range possible {
			check := lengthComp(flag+string(c), conn, r)
			if value > check {
				final = c
				value = check
				bre = 1
			}
		}
		if bre == 0 {
			break
		}
		flag = flag + string(final)
		fmt.Println("Flag Till Now: ", flag)
	}
}
