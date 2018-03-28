package main

import (
	"crypto/rand"
	"testing"
	"fmt"
	"hash"
	"github.com/hyperledger/fabric/bccsp"
)

type hasher struct {
	hash func() hash.Hash
}

func (c *hasher) Hash(msg []byte, opts bccsp.HashOpts) (hash []byte, err error) {
	h := c.hash()
	h.Write(msg)
	return h.Sum(nil), nil
}



func main() {
	/*t := &testing.T{}
	txSize := 1
	ContructRandomBytes(t, txSize)*/

	b := byte(111)
	fmt.Println(b)

	/*value := make([]byte, 2)
	_, err := rand.Read(value)
	handleErr(err)
	fmt.Println(value)
	res := hex.EncodeToString(value)
	fmt.Println(res)*/



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


func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}