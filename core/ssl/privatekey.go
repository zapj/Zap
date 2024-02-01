package ssl

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"errors"
)

func GeneratePrivateKey(keyType string, bitSize int) (crypto.Signer, error) {
	switch keyType {
	case "RSA":
		return generateRsaPrivateKey(bitSize)
	case "ECC":
		return generateEccPrivateKey(bitSize)
	}
	return nil, errors.New("KeyType not supported")
}

func generateRsaPrivateKey(bitSize int) (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bitSize)
	if err != nil {
		return nil, err
	}

	err = privateKey.Validate()
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// bitSize 224  |  256 | 384
func generateEccPrivateKey(bitSize int) (*ecdsa.PrivateKey, error) {
	var curve elliptic.Curve
	if bitSize == 256 {
		curve = elliptic.P256()
	} else if bitSize == 224 {
		curve = elliptic.P224()
	} else if bitSize == 384 {
		curve = elliptic.P384()
	}

	privateKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
