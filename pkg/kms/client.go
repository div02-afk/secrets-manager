package kms

import proto "github.com/div02-afk/secrets-manager/gen/kms"

type KMSClient interface {
	Encrypt(DEK []byte) ([]byte, error)
	Decrypt(encryptedDEK []byte) ([]byte, error)
}

type KMSService struct {
	proto.UnimplementedKMSServer

	KMS KMSClient
}