package main

import (
	"fmt"
	"ebcdic"
)

func main() {
	fmt.Println("Please enter your name")
	ASCIIName := []byte{}
	EBCDICName :=[]byte{}
	fmt.Scanln(&ASCIIName)
	fmt.Println(string(ebcdic.Encode(ASCIIName)))
	fmt.Println("Please copy and paste the output")
	fmt.Scanln(&EBCDICName)
	fmt.Println(string(ebcdic.Decode(EBCDICName)))
}
