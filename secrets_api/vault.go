package secrets_api

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)


//Vault not thread safe, to be done
type Vault struct {
	encodingKey string
	filename    string
	keyValues   map[string]string
}

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
	encryptedJson, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	decryptedJson, err := Decrypt(v.encodingKey, string(encryptedJson))
	if err != nil {
		return fmt.Errorf("while decrypting content %s, error: %v \n", string(encryptedJson), err)
	}
	err = json.Unmarshal([]byte(decryptedJson), &v.keyValues)
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
	b, err := json.Marshal(v.keyValues)
	if err != nil {
		return err
	}
	hex, err := Encrypt(v.encodingKey, string(b))
	if err != nil {
		return fmt.Errorf("While saving to file %v \n", err)
	}
	file, err := os.OpenFile(v.filename, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(hex)
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
