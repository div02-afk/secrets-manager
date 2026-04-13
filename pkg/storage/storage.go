package storage

type Storage interface {
	GetSecretValue(tenetID []byte, secretKey string,version *uint8) ([]byte, error)
	GetDEK(tenetID []byte) ([]byte, error)
	ValidateAuth(tenetID []byte, apiKey string) (bool, error)
	AddSecret(tenetID []byte, secretKey string, secretValue []byte) (error)
}
