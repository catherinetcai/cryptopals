package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

const (
	ErrInvalidHex = "Invalid hex input"
)

func main() {
	fmt.Println("vim-go")
}

func hexToBase64(input string) (string, error) {
	decoded, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(decoded), nil
}
