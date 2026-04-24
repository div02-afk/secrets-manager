package auth

type Auth interface {
	Validate(identifer int64, token string) (bool, error)
}
