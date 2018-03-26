package testutil

import (
	"testing"
	"crypto/rand"
)

func ContructRandomBytes(t testing.TB, size int) []byte {
	value := make([]byte, size)
	_, err := rand.Read(value)
	if err != nil {
		t.Fatalf("Error while generating random bytes :%s", err)
	}
	return value
}
