package pkg

type EncryptType string

type Metadata struct {
	Mode  string `json:"Mode"`
	KeyId string `json:"KeyId"`
	Data  string `json:"Data"`
}

type Codec interface {
	Marshal(value any) ([]byte, error)
	Unmarshal(data []byte, valuePtr any) error
}

type Crypto interface {
	Encrypt([]byte) ([]byte, error)
	Decrypt([]byte) ([]byte, error)
	GetType() EncryptType
}
