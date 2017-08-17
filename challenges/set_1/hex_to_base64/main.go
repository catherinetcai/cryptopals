package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
)

const (
	defaultInput = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
)

var input string

func main() {
	flag.StringVar(&input, "input", defaultInput, "Input as hex to convert to base64")
	flag.Parse()
	if input == "" {
		log.Fatal("Must provide an input!")
	}
	decoded, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal("Invalid hex: ", err.Error())
	}
	fmt.Println("Converted: ", base64.StdEncoding.EncodeToString(decoded))
}
