package storage

import "time"

type Storage interface {
	GetSecretValue(tenantID []byte, secretKey string, version *uint8) (*SecretEntry, error)
	GetDEK(tenantID []byte, DEKVersion *uint8) (*DEKDTO, error) // Set DEKVersion to nil for latest DEK entry
	ValidateAuth(tenantID []byte, apiKey string) (bool, error)
	AddSecret(tenantID []byte, secretKey string, secretValue []byte, DEKVersion uint8) (uint8, error)
}

type SecretEntry struct {
	ID          []byte
	SecretKey   string
	SecretValue []byte
	TenetID     []byte
	Version     uint8
	DEKVersion  uint8
	CreatedAt   time.Time
}
type SecretDTO struct {
	SecretKey   string `json:"secretKey"`
	SecretValue []byte `json:"secretValue"`
	TenetID     []byte `json:"tenantId"`
}

type DEKEntry struct {
	ID        []byte
	TenetID   []byte
	DEK       []byte
	Version   uint8
	CreatedAt time.Time
}
type DEKDTO struct {
	TenetID []byte
	DEK     []byte
	Version uint8
}

type AuthEntry struct {
	ID        []byte
	TenetID   []byte
	APIKey    string
	Name      string
	CreatedAt time.Time
}
