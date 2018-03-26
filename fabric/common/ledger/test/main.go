package main

import (
	"crypto/rand"
	"testing"
	"fmt"
)
func main() {
	t := &testing.T{}
	txSize := 1
	ContructRandomBytes(t, txSize)
}


func ContructRandomBytes(t testing.TB, size int) []byte {
	value := make([]byte, size)
	fmt.Println(value)
	_, err := rand.Read(value)
	fmt.Println(value)
	if err != nil {
		t.Fatalf("Error while generating random bytes :%s", err)
	}
	return value
}
