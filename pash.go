package pash

import (
	"golang.org/x/crypto/bcrypt"
)

func hash(str string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(str), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// HashMinCost uses `bcrypt.MinCost` currently `4`
func HashMinCost(str string) (string, error) {
	return hash(str, bcrypt.MinCost)
}

// Hash uses `bcrypt.DefaultCost` currently `10`
func Hash(str string) (string, error) {
	return hash(str, bcrypt.DefaultCost)
}

// HashMaxCost uses `bcrypt.MaxCost` currently `31`
func HashMaxCost(str string) (string, error) {
	return hash(str, bcrypt.MaxCost)
}

// HashWithCost accepts a variable cost factor. Use with care. The higher the cost factor, the more time it takes to
// yield a result (or crack).
func HashWithCost(str string, cost int) (string, error) {
	return hash(str, cost)
}

// VerifyHash takes the hashed string and the raw string, performs a comparison if hashes match.
// returns true if the do or false if not. Any errors will be returned as well.
func VerifyHash(hashedStr string, plainStr string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(plainStr))
	if err != nil {
		return false, err
	}
	return true, nil
}
