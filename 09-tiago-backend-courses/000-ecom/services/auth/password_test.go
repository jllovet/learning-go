package auth

import (
	"testing"
)

func TestHashedPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing the password: %v", err)
	}
	if hash == "" {
		t.Errorf("expected hash to be nonempty")
	}
	if hash == "password" {
		t.Errorf("expected hash to be different than password")
	}
}

func TestComparePasswords(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("error hashing the password: %v", err)
	}
	if !ComparePasswords(hash, "password") {
		t.Errorf("expected the hashed password to match the hash")
	}
	if ComparePasswords(hash, "notpassword") {
		t.Errorf("expected the hashed password to not match the hash")
	}
}
