package hex

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"unicode"
	"unicode/utf8"
)

const (
	defaultInput = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
)

var (
	input string
	// Letter frequency map https://www.math.cornell.edu/~mec/2003-2004/cryptography/subs/frequencies.html
	letterFrequencies = map[string]float64{
		"A": .0817,
		"B": .0150,
		"C": .0278,
		"D": .0425,
		"E": .1270,
		"F": .0223,
		"G": .0202,
		"H": .0609,
		"I": .0697,
		"J": .0015,
		"K": .0077,
		"L": .0403,
		"M": .0241,
		"N": .0675,
		"O": .0751,
		"P": .0193,
		"Q": .0010,
		"R": .0599,
		"S": .0633,
		"T": .0906,
		"U": .0276,
		"V": .0098,
		"W": .0236,
		"X": .0015,
		"Y": .0197,
		"Z": .0007,
	}
)

// Converts hex string to raw bytes
func HexToBase64(input string) []byte {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		log.Fatal("Invalid hex: ", err.Error())
	}
	return []byte(base64.StdEncoding.EncodeToString(decoded))
}

func Xor(inputA, inputB []byte) []byte {
	res := make([]byte, len(inputA))
	for i := 0; i < len(res); i++ {
		res[i] = inputA[i] ^ inputB[i]
	}
	return res
}

// Input xor'd with all possible bytes
func XorAll(input []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i <= 255; i++ {
		res := input[i] ^ byte(i)
		resRune, _ := utf8.DecodeRune([]byte{res})
		if unicode.IsLetter(resRune) {
			fmt.Printf("Letter found: %v\n", string(resRune))
			output = append(output, res)
		}
	}
	fmt.Println(string(output))
	return output
}
