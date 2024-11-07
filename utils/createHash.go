package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
)

func CreateHash(str string) (string, error) {
	if str == "" {
		return "", errors.New("input string cannot be empty")
	}

	hash := sha256.New()
	_, err := hash.Write([]byte(str))
	if err != nil {
		return "", fmt.Errorf("failed to write to hash: %w", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
