package crypto

import "crypto/rand"

const (
	NonceSize = 24
)

func GetRandomBytes(len int) ([]byte, error) {
	key := make([]byte, len)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, err
}

func GetRandomNonce() ([]byte, error) {
	return GetRandomBytes(NonceSize)
}
