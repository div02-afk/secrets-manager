package auth

import "github.com/div02-afk/secrets-manager/pkg/storage"

type APIAuth struct {
	storage storage.Storage
}

func CreateAPIAuthProvider(s storage.Storage) *APIAuth {
	return &APIAuth{
		storage: s,
	}
}

func (a *APIAuth) Validate(identifer int64,token string) (bool, error) {
	return a.storage.ValidateAuth(identifer,token)
}