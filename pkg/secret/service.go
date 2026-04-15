package secret

import (
	encryptionprovider "github.com/div02-afk/secrets-manager/pkg/encryption"
	"github.com/div02-afk/secrets-manager/pkg/kms"
	"github.com/div02-afk/secrets-manager/pkg/storage"
)

type SecretService struct {
	kms                kms.KMSClient
	storage            storage.Storage
	encryptionProvider encryptionprovider.EncrpytionProvider
}

func NewSecretService(kms kms.KMSClient, storage storage.Storage, encryptionProvider encryptionprovider.EncrpytionProvider) *SecretService {
	return &SecretService{
		kms:                kms,
		storage:            storage,
		encryptionProvider: encryptionProvider,
	}
}

func (s *SecretService) Get(tenantID []byte, secretKey string, version *uint8) ([]byte, error) {
	secretEntry, err := s.storage.GetSecretValue(tenantID, secretKey, version)
	if err != nil {
		return nil, err
	}

	decryptedDEK, _, err := s.getDecryptedDEKforTenet(tenantID, &secretEntry.DEKVersion)
	if err != nil {
		return nil, err
	}

	decryptedSecretValue, err := s.encryptionProvider.Decrypt(decryptedDEK, secretEntry.SecretValue)
	if err != nil {
		return nil, err
	}

	return decryptedSecretValue, nil
}

func (s *SecretService) Add(tenantID []byte, secretKey string, secretValue []byte) (uint8, error) {
	decryptedDEK, DEKVersion, err := s.getDecryptedDEKforTenet(tenantID, nil)
	if err != nil {
		return 0, err
	}

	encrpytedSecretValue, err := s.encryptionProvider.Encrypt(decryptedDEK, secretValue)
	if err != nil {
		return 0, err
	}

	version, err := s.storage.AddSecret(tenantID, secretKey, encrpytedSecretValue, DEKVersion)
	if err != nil {
		return 0, err
	}
	return version, nil
}

func (s *SecretService) getDecryptedDEKforTenet(tenantID []byte, DEKVersion *uint8) ([]byte, uint8, error) {
	DEKEntry, err := s.storage.GetDEK(tenantID, DEKVersion)
	if err != nil {
		return nil, 0, err
	}

	decryptedDEK, err := s.kms.Decrypt(DEKEntry.DEK)
	if err != nil {
		return nil, 0, err
	}

	return decryptedDEK, DEKEntry.Version, nil
}
