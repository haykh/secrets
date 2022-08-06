package secret

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/haykh/secrets/cipher"
)

type Vault struct {
	encodingKey string
	filepath    string
	mutex       sync.Mutex
	keyValues   map[string]string
}

func File(encodingKey, filepath string) *Vault {
	return &Vault{
		encodingKey: encodingKey,
		filepath:    filepath,
	}
}

func (v *Vault) Get(key string) (string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return "", err
	}

	value, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("secret: no value for given key")
	}
	return value, nil
}

func (v *Vault) Set(key, value string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}

	v.keyValues[key] = value
	return v.save()
}

func (v *Vault) List() ([]string, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	keys := make([]string, 0, len(v.keyValues))

	err := v.load()
	if err != nil {
		return keys, err
	}

	for key := range v.keyValues {
		keys = append(keys, key)
	}

	return keys, nil
}

func (v *Vault) Delete(key string) error {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	err := v.load()
	if err != nil {
		return err
	}

	_, ok := v.keyValues[key]
	if !ok {
		return errors.New("secret: no value for given key")
	}

	delete(v.keyValues, key)
	return v.save()
}

func (v *Vault) load() error {
	f, err := os.Open(v.filepath)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()

	r, err := cipher.DecryptReader(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.readKeyValues(r)
}

func (v *Vault) save() error {
	f, err := os.OpenFile(v.filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	w, err := cipher.EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}

	return v.writeKeyValues(w)
}

func (v *Vault) readKeyValues(r io.Reader) error {
	dec := json.NewDecoder(r)
	return dec.Decode(&v.keyValues)
}

func (v *Vault) writeKeyValues(w io.Writer) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v.keyValues)
}
