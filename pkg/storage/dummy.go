package storage

import "errors"

var errDummyNotImplemented = errors.New("dummy storage: not implemented")

type DummyStorage struct{}

var _ Storage = (*DummyStorage)(nil)

func NewDummyStorage() *DummyStorage {
	return &DummyStorage{}
}

func (d *DummyStorage) GetSecretValue(tenantID []byte, secretKey string, version *uint8) (*SecretEntry, error) {
	return nil, errDummyNotImplemented
}

func (d *DummyStorage) GetDEK(tenantID []byte, DEKVersion *uint8) (*DEKDTO, error) {
	return nil, errDummyNotImplemented
}

func (d *DummyStorage) ValidateAuth(tenantID []byte, apiKey string) (bool, error) {
	return false, errDummyNotImplemented
}

func (d *DummyStorage) AddSecret(tenantID []byte, secretKey string, secretValue []byte, DEKVersion uint8) (uint8, error) {
	return 0, errDummyNotImplemented
}
