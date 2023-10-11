package main

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"matrixAlgo"
	"strings"
)

func calculateFunc(n *big.Int) *big.Int {
	A := [][]*big.Int{
		{big.NewInt(0), big.NewInt(1), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(1), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(1)},
		{big.NewInt(55692), big.NewInt(-9549), big.NewInt(301), big.NewInt(21)},
	}
	p, _ := new(big.Int).SetString("1"+strings.Repeat("0", 10000), 10)
	matExp, _ := matrixAlgo.MatrixExponentiationModP(A, n, p)
	funValues, _ := matrixAlgo.MultiplyMatricesModP(matExp, [][]*big.Int{{big.NewInt(1)}, {big.NewInt(2)}, {big.NewInt(3)}, {big.NewInt(4)}})
	return funValues[0][0].Mod(funValues[0][0], p)
}

func decryptFlag(s string) (string, error) {
	verify := "96cc5f3b460732b442814fd33cf8537c"
	flagEnc, _ := hex.DecodeString("42cbbce1487b443de1acf4834baed794f4bbd0dfe7d7086e788af7922b")
	// Calculate MD5 hash
	md5Hash := md5.Sum([]byte(s))
	md5Hex := hex.EncodeToString(md5Hash[:])
	if md5Hex != verify {
		return "", fmt.Errorf("Mismatch of MD5. \nExpected: %s \nFound: %s", verify, md5Hex)
	}

	// Calculate SHA256 hash
	sha256Hash := sha256.Sum256([]byte(s))

	n := len(flagEnc)
	if len(sha256Hash) < n {
		n = len(sha256Hash)
	}

	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = sha256Hash[i] ^ flagEnc[i]
	}

	return string(result), nil
}

func main() {
	A := big.NewInt(20000000)
	flag, _ := decryptFlag(calculateFunc(A).String())
	fmt.Println("Flag: ", flag)
}
