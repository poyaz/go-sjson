package codec

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"reflect"

	"github.com/poyaz/go-sjson/pkg"
)

type Json struct {
	es pkg.Crypto
}

var _ pkg.Codec = &Json{}

func NewJson(encryptionService pkg.Crypto) (Json, error) {
	return Json{es: encryptionService}, nil
}

func (s Json) Marshal(value any) ([]byte, error) {
	dt := reflect.TypeOf(value)
	dv := reflect.ValueOf(value)

	res, err := s.typeEncryption(dv, dt)
	if err != nil {
		return nil, err
	}

	return json.Marshal(res)
}

func (s Json) Unmarshal(data []byte, valuePtr any) error {
	dt := reflect.TypeOf(valuePtr)
	if dt.Kind() != reflect.Pointer {
		return pkg.NewNonPointerError()
	}

	elem := dt.Elem()
	switch elem.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.String,
		reflect.Bool,
		reflect.Map:
		var dataEnc pkg.Metadata
		if err := json.Unmarshal(data, &dataEnc); err != nil {
			return err
		}

		return s.decryption([]byte(dataEnc.Data), valuePtr)
	default:
		return pkg.NewNonTypeError(dt)
	}
}

func (s Json) typeEncryption(data reflect.Value, t reflect.Type) (pkg.Metadata, error) {
	var enc []byte
	var err error

	switch t.Kind() {
	case reflect.Ptr:
		return s.typeEncryption(data.Elem(), t.Elem())
	case reflect.Map:
		return s.mapEncryption(data, t)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.String,
		reflect.Bool:
		enc, err = s.encryption(data)
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

func (s Json) mapEncryption(data reflect.Value, t reflect.Type) (pkg.Metadata, error) {
	var res pkg.Metadata
	if t.Key().Kind() == reflect.Ptr || t.Elem().Kind() == reflect.Ptr {
		return res, pkg.NewNonTypeError(t)
	}

	enc, err := s.encryption(data)
	if err != nil {
		return res, err
	}
	res = pkg.Metadata{
		Mode: "AES",
		Data: string(enc),
	}

	return res, nil
}

func (s Json) encryption(data reflect.Value) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.EncodeValue(data); err != nil {
		return nil, err
	}

	return s.es.Encrypt(buf.Bytes())
}

func (s Json) decryption(data []byte, valuePtr any) error {
	b, err := s.es.Decrypt(data)
	if err != nil {
		return err
	}

	dec := gob.NewDecoder(bytes.NewBuffer(b))

	return dec.Decode(valuePtr)
}
