package ssl

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func ExportPrivateKeyToFile(priv any, file string) error {
	pkeyBytes := ExportPrivateKeyToMemory(priv)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(pkeyBytes)
	return err
}

func ExportPrivateKeyToMemory(priv any) []byte {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		pKeyPEM := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(k),
		}
		return pem.EncodeToMemory(pKeyPEM)
	case *ecdsa.PrivateKey:
		pkeyBytes, _ := x509.MarshalECPrivateKey(k)
		pKeyPEM := &pem.Block{
			Type:  "EC PRIVATE KEY",
			Bytes: pkeyBytes,
		}
		return pem.EncodeToMemory(pKeyPEM)
	case ed25519.PrivateKey:
		pkeyBytes, _ := x509.MarshalPKCS8PrivateKey(k)
		pKeyPEM := &pem.Block{
			Type:  "PRIVATE KEY",
			Bytes: pkeyBytes,
		}
		return pem.EncodeToMemory(pKeyPEM)
	default:
		return nil
	}

}

func ExportPublicKeyToFile(priv any, file string) error {
	pubBytes := ExportPublicKeyToMemory(priv)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(pubBytes)
	return err
}

func ExportPublicKeyToMemory(priv any) []byte {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		publicKeyPEM := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(&k.PublicKey),
		}
		return pem.EncodeToMemory(publicKeyPEM)
	case *ecdsa.PrivateKey:
		pubBytes, _ := x509.MarshalPKIXPublicKey(&k.PublicKey)
		publicKeyPEM := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubBytes,
		}
		return pem.EncodeToMemory(publicKeyPEM)
	case ed25519.PrivateKey:
		pubBytes, _ := x509.MarshalPKIXPublicKey(k.Public().(ed25519.PublicKey))
		publicKeyPEM := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubBytes,
		}
		return pem.EncodeToMemory(publicKeyPEM)
	default:
		return nil
	}

}
