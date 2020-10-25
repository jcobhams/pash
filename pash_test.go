package pash

import (
	"testing"
)

func TestHashAndVerify(t *testing.T) {
	samplePassword := "MYsUp3rStr0ng"
	hash, err := Hash(samplePassword)

	if hash == "" {
		t.Errorf("Hash should not be empty")
	}

	if err != nil {
		t.Errorf("Expected error to be nil")
	}

	ok, err := VerifyHash(hash, samplePassword)
	if !ok {
		t.Errorf("hash verification should pass")
	}
	if err != nil {
		t.Errorf("Expected error to be nil")
	}
}
