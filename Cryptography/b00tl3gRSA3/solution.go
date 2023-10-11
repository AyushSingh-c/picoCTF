package main

import (
	"factordbapi"
	"fmt"
	"math/big"
	"rsaDecrypt"
)

func EularTotient(factors []factordbapi.Factor) *big.Int {
	var (
		temp1 = big.NewInt(0)
		temp2 = big.NewInt(0)
		prod  = big.NewInt(1)
	)
	for _, i := range factors {
		n, _ := new(big.Int).SetString(i.Number, 10)
		e := big.NewInt(int64(i.Power))
		temp1.Exp(n, e, nil)
		temp2.Exp(n, e.Sub(e, big.NewInt(1)), nil)
		temp1.Sub(temp1, temp2)
		prod.Mul(prod, temp1)
	}
	return prod
}

func DecryptRSA_MultiplePQ(rsaInfo rsaDecrypt.PublicInfo) *big.Int {
	var (
		temp1 = big.NewInt(0)
	)

	factors, _ := factordbapi.GetFactors(rsaInfo.N.String())
	phi := EularTotient(factors)
	phi, _ = new(big.Int).SetString("14168486611611863348973674556400046363159171601110968706493641768051277568817946395257419042951484539682067055152658940205392717234238702330918717888787027840957700606066422422773313796148026418846406573103555657973288397389429093428423196263958230383013451184145339902509597469127227141604911090835317284830412672901390788622752029868032000000", 10)
	d, _ := rsaDecrypt.ModInverse(rsaInfo.E, phi)
	fmt.Println("d: ", d)
	fmt.Println("e: ", rsaInfo.E)
	fmt.Println("2^d: ", temp1.Exp(big.NewInt(2), d, rsaInfo.N))
	fmt.Println("2^ed: ", temp1.Exp(temp1, rsaInfo.E, rsaInfo.N))
	fmt.Println("2^e: ", temp1.Exp(big.NewInt(2), rsaInfo.E, rsaInfo.N))
	m := new(big.Int).Exp(rsaInfo.C, d, rsaInfo.N)

	return m
}

func main() {
	n, _ := new(big.Int).SetString("14168486651162826595731488311843536847403928582001778545516796264413940099735159800564298641411168786861862940211743208011537868324584319718982459821738030729819163724216517306077067467359204271707846009503384758574176159161030342724383957197821685256062956112255459126699683657914864059215895529047556231397975929282241859269161073496554442177", 10)
	c, _ := new(big.Int).SetString("4111881688187565544791087261917401152624964586782423676244769056281317716265124587751577910502348509808630788972842365376358261763672569270644149841608772539939469871126066913393046028812108214426730356295585739995690469691496127836743592713753235956408068301399281149209423081819268309501739358904588190480003659708421658189242825886599848844", 10)
	e, _ := new(big.Int).SetString("65537", 10)
	factors, _ := factordbapi.GetFactors(n.String())
	phi := EularTotient(factors)
	fmt.Println(factors)
	fmt.Println("phi: ", phi)
	fmt.Printf("Flag: %s", DecryptRSA_MultiplePQ(rsaDecrypt.PublicInfo{N: n, C: c, E: e}).Bytes())
}
