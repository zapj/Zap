package ssl

func GenerateSSHKeyPairs(keyType string) {
	bitSize := 256
	if keyType == "RSA" {
		bitSize = 2048
	} else if keyType == "EC" {
		bitSize = 256
	}
	_, _ = GeneratePrivateKey(keyType, bitSize)

}

// pubKey, privKey, _ := ed25519.GenerateKey(rand.Reader)
// publicKey, _ := ssh.NewPublicKey(pubKey)

// pemKey := &pem.Block{
//     Type:  "OPENSSH PRIVATE KEY",
//     Bytes: MarshalED25519PrivateKey(privKey)
// }
// privateKey := pem.EncodeToMemory(pemKey)
// authorizedKey := ssh.MarshalAuthorizedKey(publicKey)

// _ = ioutil.WriteFile("id_ed25519", privateKey, 0600)
// _ = ioutil.WriteFile("id_ed25519.pub", authorizedKey, 0644)
