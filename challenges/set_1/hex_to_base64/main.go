package main

import (
	"encoding/base64"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
)

var input string

func main() {
	flag.StringVar(&input, "input", "", "Input as hex to convert to base64")
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
