package auth

import (
	"strings"
	"testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "correct horse battery staple"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	if !strings.HasPrefix(hash, "$argon2id$") {
		t.Fatalf("unexpected hash format: %s", hash)
	}

	valid, err := VerifyPassword(hash, password)
	if err != nil {
		t.Fatalf("VerifyPassword returned an error: %v", err)
	}

	if !valid {
		t.Fatal("correct password was rejected")
	}
}

func TestVerifyWrongPassword(t *testing.T) {
	hash, err := HashPassword("correct password")
	if err != nil {
		t.Fatalf("HashPassword returned an error: %v", err)
	}

	valid, err := VerifyPassword(hash, "wrong password")
	if err != nil {
		t.Fatalf("VerifyPassword returned an error: %v", err)
	}

	if valid {
		t.Fatal("incorrect password was accepted")
	}
}

func TestHashesUseDifferentSalts(t *testing.T) {
	password := "same password"

	hash1, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}

	hash2, err := HashPassword(password)
	if err != nil {
		t.Fatal(err)
	}

	if hash1 == hash2 {
		t.Fatal("identical passwords produced identical hashes")
	}
}

func TestRejectMalformedHash(t *testing.T) {
	_, err := VerifyPassword("not a valid hash", "password")
	if err == nil {
		t.Fatal("expected an error for malformed hash")
	}
}

func TestRejectPasswordTooLong(t *testing.T) {
	password := strings.Repeat("a", maxPasswordBytes+1)

	_, err := HashPassword(password)
	if err != ErrPasswordTooLong {
		t.Fatalf("expected ErrPasswordTooLong, got %v", err)
	}
}
