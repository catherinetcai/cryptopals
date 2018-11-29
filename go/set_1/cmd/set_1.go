package cmd

import (
	"encoding/hex"
	"fmt"

	setone "github.com/catherinetcai/cryptopals/set_1/hex"
	"github.com/spf13/cobra"
)

var (
	input string
)

var SetOneCmd = &cobra.Command{
	Use:   "setOne",
	Short: "Set one",
}

var partOneCmd = &cobra.Command{
	Use: "one",
	Run: partOne,
}

var partTwoCmd = &cobra.Command{
	Use: "two",
	Run: partTwo,
}

var partThreeCmd = &cobra.Command{
	Use: "three",
	Run: partThree,
}

func init() {
	RootCmd.AddCommand(SetOneCmd)
	SetOneCmd.AddCommand(partOneCmd)
	SetOneCmd.AddCommand(partTwoCmd)
	SetOneCmd.AddCommand(partThreeCmd)
	SetOneCmd.PersistentFlags().StringVarP(&input, "input", "i", "", "Input")
}

func partOne(cmd *cobra.Command, args []string) {
	input = "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	fmt.Println(string(setone.HexToBase64(input)))
}

func partTwo(cmd *cobra.Command, args []string) {
	inputA := "1c0111001f010100061a024b53535009181c"
	inputB := "686974207468652062756c6c277320657965"

	bytesA, _ := hex.DecodeString(inputA)
	bytesB, _ := hex.DecodeString(inputB)

	res := setone.Xor(bytesA, bytesB)

	fmt.Println(hex.EncodeToString(res))
}

func partThree(cmd *cobra.Command, args []string) {
	input = "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	inputBytes, _ := hex.DecodeString(input)
	setone.XorAll(inputBytes)
}
