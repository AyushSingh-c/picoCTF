package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"rsaDecrypt"
	"strings"
)

func readFile(file string) (n, c, e []*big.Int) {
	// open the file for reading
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value, _ := new(big.Int).SetString(strings.TrimSpace(parts[1]), 10)

		switch key {
		case "n":
			n = append(n, value)
		case "c":
			c = append(c, value)
		case "e":
			e = append(e, value)
		default:
			log.Fatal("unknown key")
		}
	}
	// check for any error from scanning
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return n, c, e
}

func main() {
	N, C, E := readFile("encrypted-messages.txt")

	Messages := []string{"I just cannot wait for rowing practice today!",
		"I hope we win that big rowing match next week!",
		"Rowing is such a fun sport!"}

	FlagInfo := []rsaDecrypt.PublicInfo{}
	for i := 0; i < len(N); i++ {
		check := false
		rsaInfo := rsaDecrypt.PublicInfo{N: N[i], C: C[i], E: E[i]}
		for _, message := range Messages {
			check = check || rsaDecrypt.CheckDecrypt(rsaInfo, message)
		}
		if !check {
			FlagInfo = append(FlagInfo, rsaInfo)
		}
	}
	flag_power, _ := rsaDecrypt.CrtAttack(FlagInfo)
	flag, _ := rsaDecrypt.Kthroot_halley(big.NewInt(3), flag_power)
	fmt.Printf("Decrypted Plain Text: %s\n", flag.Bytes())
}
