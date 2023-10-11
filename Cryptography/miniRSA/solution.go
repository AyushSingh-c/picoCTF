package main

import (
	"fmt"
	"formatCTF"
	"math/big"
	"rsaDecrypt"
	"strings"
)

func main() {
	n, _ := new(big.Int).SetString("29331922499794985782735976045591164936683059380558950386560160105740343201513369939006307531165922708949619162698623675349030430859547825708994708321803705309459438099340427770580064400911431856656901982789948285309956111848686906152664473350940486507451771223435835260168971210087470894448460745593956840586530527915802541450092946574694809584880896601317519794442862977471129319781313161842056501715040555964011899589002863730868679527184420789010551475067862907739054966183120621407246398518098981106431219207697870293412176440482900183550467375190239898455201170831410460483829448603477361305838743852756938687673", 10)
	c, _ := new(big.Int).SetString("2205316413931134031074603746928247799030155221252519872649649212867614751848436763801274360463406171277838056821437115883619169702963504606017565783537203207707757768473109845162808575425972525116337319108047893250549462147185741761825125", 10)
	e, _ := new(big.Int).SetString("3", 10)

	find_flag := func(i int) {
		test := big.NewInt(0)
		test.Mul(big.NewInt(int64(i)), n)
		plainText, _ := rsaDecrypt.Kthroot_halley(e, test.Add(test, c))
		flag := string(plainText.Bytes())
		if formatCTF.IsStringPrintable(flag) {
			fmt.Printf("Decrypted Plain Text: %s\n", strings.TrimSpace(flag))
		}
	}
	formatCTF.RunParallel(find_flag, 10000)
}
