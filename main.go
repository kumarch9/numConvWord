package main

import (
	"fmt"
	"log"

	"github.com/pkchandra/numconvword/convword"
)

func main() {
	var num int
	fmt.Println("Enter any number and the number must be less than 99 crores.")
	fmt.Scanf("%d", &num)
	fmt.Printf("You entered the  number is  %+v\n", num)
	word, err := convword.ConvNumToWords(num)
	if err != nil {
		log.Fatal("err in ConvNumToWords : ", err)
	}
	fmt.Printf("Output of number = %s", word)
}
