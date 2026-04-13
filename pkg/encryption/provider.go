package encryptionprovider

type EncrpytionProvider interface {
	Encrypt(key []byte,payload []byte) ([]byte,error)
	Decrypt(key []byte,payload []byte) ([]byte,error)
}