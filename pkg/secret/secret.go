package secret

import (
	encryptionprovider "github.com/div02-afk/secrets-manager/pkg/encryption-provider"
	"github.com/div02-afk/secrets-manager/pkg/kms"
	"github.com/div02-afk/secrets-manager/pkg/storage"
)

type SecretService struct {
	kms                *kms.KMS
	storage            storage.Storage
	encryptionprovider encryptionprovider.EncrpytionProvider
}

func (s *SecretService) Get(tenetID []byte, secretKey string, version *uint8) ([]byte, error) {
	encryptedSecretValue, err := s.storage.GetSecretValue(tenetID, secretKey, version)
	if err != nil {
		return nil, err
	}

	decryptedDEK, err := s.getDecryptedDEKforTenet(tenetID)
	if err != nil {
		return nil, err
	}

	decryptedSecretValue, err := s.encryptionprovider.Decrypt(decryptedDEK, encryptedSecretValue)
	if err != nil {
		return nil, err
	}

	return decryptedSecretValue, nil
}

func (s *SecretService) Add(tenetID []byte, secretKey string, secretValue []byte) error {
	decryptedDEK, err := s.getDecryptedDEKforTenet(tenetID)
	if err != nil {
		return err
	}

	encrpytedSecretValue, err := s.encryptionprovider.Encrypt(decryptedDEK, secretValue)
	if err != nil {
		return err
	}

	err = s.storage.AddSecret(tenetID, secretKey, encrpytedSecretValue)

	return err
}

func (s *SecretService) getDecryptedDEKforTenet(tenetID []byte) ([]byte, error) {
	encryptedDEK, err := s.storage.GetDEK(tenetID)
	if err != nil {
		return nil, err
	}

	decryptedDEK, err := s.kms.Decrypt(encryptedDEK)
	if err != nil {
		return nil, err
	}

	return decryptedDEK, nil
}