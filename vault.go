package secret

import (
	"errors"

	"github.com/patelevans/secret/encrypt"
)

type Vault struct {
	encodingKey string
	keyVaules   map[string]string
}

func Memory(encodingKey string) Vault {
	return Vault{
		encodingKey: encodingKey,
		keyVaules:   make(map[string]string),
	}
}

func (v *Vault) Get(key string) (string, error) {
	hex, ok := v.keyVaules[key]
	if !ok {
		return "", errors.New("secret: no value for given key")
	}

	decryptedValue, err := encrypt.Decrypt(v.encodingKey, hex)
	if err != nil {
		return "", err
	}
	return decryptedValue, nil
}

func (v *Vault) Set(key, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}

	v.keyVaules[key] = encryptedValue
	return nil
}
