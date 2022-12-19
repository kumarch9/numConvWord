package convword

import (
	"log"
	"strconv"
	"strings"
)

var (
	Num0ToWord = map[int]string{
		0: "Zero", 1: "One", 2: "Two", 3: "three", 4: "Four", 5: "Five", 6: "Six", 7: "Seven",
		8: "Eight", 9: "Nine", 10: "Ten", 11: "Eleven", 12: "Twelve", 13: "Thirteen", 14: "Fourteen",
		15: "Fifteen", 16: "Sixteen", 17: "Seventeen", 18: "Eighteen", 19: "Nineteen", 20: "Twenty",
		30: "Thirty", 40: "Forty", 50: "Fifty", 60: "Sixty", 70: "Seventy", 80: "Eighty", 90: "Ninety",
	}

	NumCompose = map[int]string{
		1: "", 2: "", 3: "Hundreds", 4: "Thousands", 5: "Thousands", 6: "Lacs", 7: "Lacs", 8: "Crore",
		9: "Crore",
	}
)

func Convert0To99(n int) string {
	s := []string{}
	var str string
	if n < 20 {
		w := Num0ToWord[n]
		s = append(s, w)
		str = strings.Join(s, " ")
		return str
	}
	if n > 19 && n < 100 {
		preNum := n / 10
		postNum := n % 10
		preWord := preNum * 10
		if postNum == 0 {
			s = append(s, Num0ToWord[n])
			str = strings.Join(s, " ")
			return str
		}
		s = append(s, Num0ToWord[preWord])
		s = append(s, Num0ToWord[postNum])
	}
	str = strings.Join(s, " ")
	return str
}

/*
PowInInt() function use and it will take argument in int data type number  and return in power.
e.g. PowInInt(2,3) output become  8 number.
*/
func PowInInt(n, pow int) int {
	// nStr := strconv.Itoa(pow)
	// nLen := len(nStr)
	val := 1
	if pow == 1 {
		return val * n
	}
	for i := 0; i < pow; i++ {
		val = val * n
	}
	return val
}

/*
ConvNumToWords() function need  maximum 9 digit as e.g. 999999999 and
it will convert digit into words as per need of developer and also may
modify it.

	e.g. ConvNumToWords(8455) output will come in words(Eight Thousands Four Hundreds Fifty Five).
*/
func ConvNumToWords(number int) (string, error) {
	num := number
	s := []string{}
	var powerOfNum int
	var wordsOfNum string
label:
	numStr := strconv.Itoa(num)
	num_len := len(numStr)
	if num_len > 9 {
		log.Fatalln("number value is greater than 9 digits : ", num_len)
	}
	if num_len > 2 {
		pair := num_len % 2
		if num_len > 4 {
			if pair == 0 {
				powerOfNum = PowInInt(10, num_len-1)
			} else {
				powerOfNum = PowInInt(10, num_len-2)
			}
		} else {
			powerOfNum = PowInInt(10, num_len-1)
		}
		preNum := num / powerOfNum
		postNum := num % powerOfNum
		if postNum == 0 {
			s = append(s, Convert0To99(preNum))
			s = append(s, NumCompose[num_len])
			goto saveWord
			// wordsOfNum = strings.Join(s, " ")
			// return wordsOfNum, nil
		}
		s = append(s, Convert0To99(preNum))
		s = append(s, NumCompose[num_len])
		if postNum < 100 {
			s = append(s, Convert0To99(postNum))
			goto saveWord
			// wordsOfNum = strings.Join(s, " ")
			// return wordsOfNum, nil
		}
		num = postNum
		goto label
	}
	s = append(s, Convert0To99(num))
saveWord:
	wordsOfNum = strings.Join(s, " ")
	return wordsOfNum, nil
}
