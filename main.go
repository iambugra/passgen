package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"math/big"
	"os"
)

const (
	lowercase = "abcdefghijklmnopqrstuvwxyz"
	uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits    = "0123456789"
	symbols   = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

func GeneratePassword(length int, includeUpper, includeDigits, includeSymbols bool) (string, error) {
	if length < 1 {
		return "", fmt.Errorf("password length must be at least 1")
	}

	// Build character set
	charset := lowercase
	if includeUpper {
		charset += uppercase
	}
	if includeDigits {
		charset += digits
	}
	if includeSymbols {
		charset += symbols
	}

	password := make([]byte, length)
	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}

	return string(password), nil
}

func main() {
	// Command-line flags
	length := flag.Int("length", 16, "Length of the password")
	includeUpper := flag.Bool("upper", true, "Include uppercase letters")
	includeDigits := flag.Bool("digits", true, "Include digits")
	includeSymbols := flag.Bool("symbols", true, "Include symbols")

	flag.Parse()

	if *length < 1 {
		fmt.Println("Error: password length must be at least 1")
		os.Exit(1)
	}

	password, err := GeneratePassword(*length, *includeUpper, *includeDigits, *includeSymbols)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(password)
}
