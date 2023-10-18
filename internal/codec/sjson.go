package codec

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"reflect"

	pkg "github.com/poyaz/go-sjson/pkg"
)

type Json struct {
	es pkg.Crypto
}

var _ pkg.Codec = &Json{}

func NewJson(encryptionService pkg.Crypto) (*Json, error) {
	return &Json{es: encryptionService}, nil
}

func (s *Json) Marshal(value any) ([]byte, error) {
	dt := reflect.TypeOf(value)
	dv := reflect.ValueOf(value)

	res, err := s.typeEncryption(dv, dt)
	if err != nil {
		return nil, err
	}

	return json.Marshal(res)
}

func (s *Json) Unmarshal(data []byte, valuePtr any) error {
	dt := reflect.TypeOf(valuePtr)
	if dt.Kind() != reflect.Pointer {
		return pkg.NewNonPointerError()
	}

	switch dt.Kind() {
	default:
		var dataEnc pkg.Metadata
		if err := json.Unmarshal(data, &dataEnc); err != nil {
			return err
		}

		return s.byteDecryption([]byte(dataEnc.Data), valuePtr)
	}
}

func (s *Json) typeEncryption(data reflect.Value, t reflect.Type) (pkg.Metadata, error) {
	var enc []byte
	var err error

	switch t.Kind() {
	case reflect.Ptr:
		return s.typeEncryption(data.Elem(), t.Elem())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.String,
		reflect.Bool:
		enc, err = s.byteEncryption(data)
	default:
		return pkg.Metadata{}, pkg.NewNonTypeError(t)
	}

	var res pkg.Metadata
	if err != nil {
		return res, err
	}
	res = pkg.Metadata{
		Mode: "AES",
		Data: string(enc),
	}

	return res, nil
}

func (s *Json) byteEncryption(data reflect.Value) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.EncodeValue(data); err != nil {
		return nil, err
	}

	return s.es.Encrypt(buf.Bytes())
}

func (s *Json) byteDecryption(data []byte, valuePtr any) error {
	b, err := s.es.Decrypt(data)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(bytes.NewBuffer(b))

	return dec.Decode(valuePtr)
}
