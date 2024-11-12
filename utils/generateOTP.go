package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP(maxDigits uint32) (string, error) {
	if maxDigits == 0 {
		return "", fmt.Errorf("maxDigits must be greater than 0")
	}

	upperLimit := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(maxDigits)), nil)

	bi, err := rand.Int(rand.Reader, upperLimit)
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}

	otp := fmt.Sprintf("%0*d", maxDigits, bi)
	return otp, nil
}
