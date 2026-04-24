package kms

import (
	"os"

	encryptionprovider "github.com/div02-afk/secrets-manager/pkg/encryption"
)

type KMS struct {
	masterKey          []byte
	encryptionProvider encryptionprovider.EncrpytionProvider
}

func CreateKMSProvider(e encryptionprovider.EncrpytionProvider) *KMS{
	return &KMS{
		masterKey: []byte(os.Getenv("MASTER_KEY")),
		encryptionProvider: e,
	}
}



func (k *KMS) Encrypt(DEK []byte) ([]byte, error) {
	// TODO: implement encryption logic
	cipherTextDEK, err := k.encryptionProvider.Encrypt(k.masterKey, DEK)
	if err != nil {
		return nil, err
	}

	return cipherTextDEK, nil
}

func (k *KMS) Decrypt(encryptedDEK []byte) ([]byte, error) {
	plainTextDEK, err := k.encryptionProvider.Decrypt(k.masterKey, encryptedDEK)
	if err != nil {
		return nil, err
	}

	return plainTextDEK, nil
}
