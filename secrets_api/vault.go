package secrets_api

import (
	"context"
	"encoding/json"
	"os"
)

// Vault not thread safe, to be done
type Vault struct {
	encodingKey string
	filename    string
	keyValues   map[string]string
}

var ctx context.Context

func NewVault(encodingKey, filename string) *Vault {
	return &Vault{encodingKey: encodingKey, filename: filename}
}

func (v *Vault) loadFromFile() error {
	f, err := os.Open(v.filename)
	if err != nil {
		v.keyValues = make(map[string]string)
		return nil
	}
	defer f.Close()
	//file -> decryptor -> json
	decryptedFileReader, err := DecryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(decryptedFileReader)
	err = dec.Decode(&v.keyValues)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) Get(key string) (string, error) {
	err := v.loadFromFile()
	if err != nil {
		return "", err
	}
	return v.keyValues[key], nil
}

func (v *Vault) saveToFile() error {
	//open file, cr
	f, err := os.OpenFile(v.filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()
	encryptedFileWriter, err := EncryptWriter(v.encodingKey, f)
	if err != nil {
		return err
	}
	enc := json.NewEncoder(encryptedFileWriter)
	err = enc.Encode(v.keyValues)
	if err != nil {
		return err
	}
	return nil
}

func (v *Vault) Add(key, value string) error {
	err := v.loadFromFile()
	if err != nil {
		return err
	}
	v.keyValues[key] = value
	err = v.saveToFile()
	if err != nil {
		return err
	}
	return nil
}
