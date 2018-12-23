package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"
)

const (
	defaultInput = "1c0111001f010100061a024b53535009181c"
	defaultXor   = "686974207468652062756c6c277320657965"
)

var input string
var xorInput string

func main() {
	flag.StringVar(&input, "input", defaultInput, "Input as hex to convert to base64")
	flag.StringVar(&xorInput, "xor", defaultXor, "Input to xor against the hex input")
	flag.Parse()
	if len(input) != len(xorInput) {
		log.Fatal("Input and xorInput must be the same length")
	}
	decoded, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal("Invalid hex: ", err.Error())
	}
	xorByte := []byte(xorInput)
	fmt.Println(string(xor(decoded, xorByte)))
}

func xor(input, xorInput []byte) []byte {
	output := make([]byte, len(input))
	for i := range input {
		output[i] = input[i] ^ xorInput[i]
	}
	return output
}
