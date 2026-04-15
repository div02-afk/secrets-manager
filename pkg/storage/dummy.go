package storage

import "errors"

var errDummyNotImplemented = errors.New("dummy storage: not implemented")

type DummyStorage struct{}

var _ Storage = (*DummyStorage)(nil)

func NewDummyStorage() *DummyStorage {
	return &DummyStorage{}
}

func (d *DummyStorage) GetSecretValue(tenantID int64, secretKey string, version *int) (*SecretEntry, error) {
	return nil, errDummyNotImplemented
}

func (d *DummyStorage) GetDEK(tenantID int64, DEKVersion *int) (*DEKDTO, error) {
	return nil, errDummyNotImplemented
}

func (d *DummyStorage) ValidateAuth(tenantID int64, apiKey string) (bool, error) {
	return false, errDummyNotImplemented
}

func (d *DummyStorage) AddSecret(tenantID int64, secretKey string, secretValue []byte, DEKVersion int) (int, error) {
	return 0, errDummyNotImplemented
}
