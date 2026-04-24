package kms

import (
	"context"

	proto "github.com/div02-afk/secrets-manager/gen/kms"
)

type KMSClient interface {
	Encrypt(DEK []byte) ([]byte, error)
	Decrypt(encryptedDEK []byte) ([]byte, error)
}

type KMSService struct {
	proto.UnimplementedKMSServer

	KMS KMSClient
}

func (s KMSService) Encrypt(ctx context.Context, req *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	encrypted, err := s.KMS.Encrypt(req.Dek)
	if err != nil {
		return nil, err
	}
	return &proto.EncryptResponse{EncryptedDek: encrypted}, nil
}

func (s KMSService) Decrypt(ctx context.Context, req *proto.DecryptRequest) (*proto.DecryptResponse, error) {
	dek, err := s.KMS.Decrypt(req.EncryptedDek)
	if err != nil {
		return nil, err
	}
	return &proto.DecryptResponse{Dek: dek}, nil
}
