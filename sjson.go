package go_sjson

type Sjson struct {
	es Crypto
}

var _ Codec = &Sjson{}

func New(encryptionService Crypto) *Sjson {
	return &Sjson{es: encryptionService}
}

func (s Sjson) Marshal(value any) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (s Sjson) Unmarshal(data []byte, valuePtr any) error {
	//TODO implement me
	panic("implement me")
}
