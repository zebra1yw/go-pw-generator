package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate random passwords",
	Long: `Generate random passwords with customizable options.
for example:

pw generate -l 8 -u -d -s`,

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntP("length", "l", 8, "length of the password")
	generateCmd.Flags().BoolP("uppercase", "u", false, "Include uppercase alphabets in the password")
	generateCmd.Flags().BoolP("numbers", "d", false, "Include numbers in the password")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include symbols in the password")
}

func generatePassword(cmd *cobra.Command, args []string) {
	length, _ := cmd.Flags().GetInt("length")
	isUppercase, _ := cmd.Flags().GetBool("uppercase")
	isNumber, _ := cmd.Flags().GetBool("numbers")
	isSymbol, _ := cmd.Flags().GetBool("symbols")

	charset := "abcdefghijklmnopqrstuvwxyz"

	if isUppercase {
		charset += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	if isNumber {
		charset += "0123456789"
	}

	if isSymbol {
		charset += "!@#$%&*()-_=+[]{},;.<>?|"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	fmt.Printf("your generated password: \n %v", string(password))
}
