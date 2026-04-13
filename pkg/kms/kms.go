package kms

import encryptionprovider "github.com/div02-afk/secrets-manager/pkg/encryption"

type KMS struct {
	masterKey          []byte
	encryptionprovider encryptionprovider.EncrpytionProvider
}

func (k *KMS) Encrypt(DEK []byte) ([]byte, error) {
	// TODO: implement encryption logic
	return nil, nil
}

func (k *KMS) Decrypt(encryptedDEK []byte) ([]byte, error) {
	//TODO: implement decryption logic
	return nil, nil
}