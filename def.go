package go_sjson

type EncryptType string

type Codec interface {
	Marshal(value any) ([]byte, error)
	Unmarshal(data []byte, valuePtr any) error
}

type Crypto interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
	GetType() EncryptType
}
