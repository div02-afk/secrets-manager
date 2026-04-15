package storage

import "time"

type Storage interface {
	GetSecretValue(tenantID int64, secretKey string, version *int) (*SecretEntry, error)
	GetDEK(tenantID int64, DEKVersion *int) (*DEKDTO, error) // Set DEKVersion to nil for latest DEK entry
	ValidateAuth(tenantID int64, apiKey string) (bool, error)
	AddSecret(tenantID int64, secretKey string, secretValue []byte, DEKVersion int) (int, error)
}

type TenantEntry struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type TenantDTO struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

type SecretEntry struct {
	ID          int64     `json:"id"`
	SecretKey   string    `json:"secretKey"`
	SecretValue []byte    `json:"secretValue"`
	TenantID    int64     `json:"tenantId"`
	Version     int       `json:"version"`
	DEKVersion  int       `json:"dekVersion"`
	CreatedAt   time.Time `json:"createdAt"`
}

type SecretDTO struct {
	SecretKey   string `json:"secretKey"`
	SecretValue []byte `json:"secretValue"`
	TenantID    int64  `json:"tenantId"`
}

type DEKEntry struct {
	ID        int64     `json:"id"`
	TenantID  int64     `json:"tenantId"`
	DEK       []byte    `json:"dek"`
	Version   int       `json:"version"`
	CreatedAt time.Time `json:"createdAt"`
}

type DEKDTO struct {
	TenantID int64  `json:"tenantId"`
	DEK      []byte `json:"dek"`
	Version  int    `json:"version"`
}

type AuthEntry struct {
	ID        int64     `json:"id"`
	TenantID  int64     `json:"tenantId"`
	APIKey    string    `json:"apiKey"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuthDTO struct {
	TenantID int64  `json:"tenantId"`
	APIKey   string `json:"apiKey"`
}
