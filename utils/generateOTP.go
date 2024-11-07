package utils

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func GenerateOTP(maxDigits uint32) (string, error) {
	if maxDigits == 0 {
		return "", fmt.Errorf("maxDigits must be greater than 0")
	}

	upperLimit := big.NewInt(int64(math.Pow(10, float64(maxDigits))))

	bi, err := rand.Int(rand.Reader, upperLimit)
	if err != nil {
		return "", fmt.Errorf("failed to generate OTP: %w", err)
	}

	otp := fmt.Sprintf("%0*d", maxDigits, bi)
	return otp, nil
}
