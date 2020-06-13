package util

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

func EncodeS8Object(v interface{}) (string, error) {
	var buf bytes.Buffer

	enc := yaml.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {

		return "", err
	}
	return buf.String(), nil
}

func DecodeS8Object(value []byte, o interface{}) error {
	err := yaml.Unmarshal(value, o)
	if err != nil {
		return err
	}
	return nil
}
