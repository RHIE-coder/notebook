package demo

import (
	"crypto/rand"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateSecret(byteLength int) ([]byte, error) {
	randomSecretData := make([]byte, byteLength)
	_, err := rand.Read(randomSecretData)
	if err != nil {
		return nil, fmt.Errorf("fail to generate secret key")
	}
	return randomSecretData, nil
}

func GetSecretHashFrom(secretKey []byte) string {
	return crypto.Keccak256Hash(secretKey).Hex()
}

func GetSecretHashByte32From(secretHash string) ([32]byte, error) {
	var secretHashBytes [32]byte
	if len(secretHash) != 64 {
		var emptyHash [32]byte
		return emptyHash, fmt.Errorf("secretHash should be a 32-byte hexadecimal string") //cannot use nil as [32]byte value in return statement
	}
	copy(secretHashBytes[:], common.Hex2Bytes(secretHash))
	return secretHashBytes, nil
}
