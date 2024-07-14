package simpleEncryptFile

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"

	"filippo.io/age"
)

var (
	ErrKeyAlreadyExists = fmt.Errorf("Key already exists")
)

func GenerateKeys(path string, password string) error {
	x25519, err := age.GenerateX25519Identity()
	if err != nil {
		return fmt.Errorf("failed to generate key %w", err)
	}

	var output io.WriteCloser

	_, err = os.Stat(path)
	if err == nil {
		return ErrKeyAlreadyExists
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to generate key %w", err)
	}

	if password != "" {
		scryptRecipient, err := age.NewScryptRecipient(password)
		if err != nil {
			return fmt.Errorf("failed to generate key %w", err)
		}

		output, err = age.Encrypt(file, scryptRecipient)
		if err != nil {
			return fmt.Errorf("failed to generate key %w", err)
		}
	} else {
		output = file
	}

	if _, err := io.WriteString(output, x25519.String()); err != nil {
		return fmt.Errorf("failed to generate key %w", err)
	}

	if err := output.Close(); err != nil {
		return fmt.Errorf("failed to generate key %w", err)
	}

	if err := file.Close(); err != nil && !errors.Is(err, os.ErrClosed) {
		return fmt.Errorf("failed to generate key %w", err)
	}

	x25519pub := x25519.Recipient()

	file, err = os.Create(path + ".pub")
	if err != nil {
		return fmt.Errorf("failed to generate public key %w", err)
	}

	if _, err := io.WriteString(file, x25519pub.String()); err != nil {
		return fmt.Errorf("failed to generate public key %w", err)
	}

	if err := file.Close(); err != nil {
		return fmt.Errorf("failed to generate public key %w", err)
	}

	return nil
}

func LoadKey(path string, password string) (*age.X25519Identity, error) {
	if path == "" {
		return nil, fmt.Errorf("Failed to load key, empty path")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to load key %w", err)
	}
	defer file.Close()

	var keyReader io.Reader

	if password != "" {
		identity, err := age.NewScryptIdentity(password)
		if err != nil {
			return nil, fmt.Errorf("Failed to load key %w", err)
		}
		keyReader, err = age.Decrypt(file, identity)
		if err != nil {
			return nil, fmt.Errorf("Failed to load key %w", err)
		}
	} else {
		keyReader = file
	}

	if b, err := io.ReadAll(keyReader); err != nil {
		return nil, fmt.Errorf("Failed to load key %w", err)
	} else {
		ret, err := age.ParseX25519Identity(string(b))
		if err != nil {
			return nil, fmt.Errorf("Failed to load key %w", err)
		}
		return ret, nil
	}
}

func LoadPubKey(path string) (*age.X25519Recipient, error) {
	if path == "" {
		return nil, fmt.Errorf("Failed to load public key, empty path")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to load Public key %w", err)
	}
	keyByte, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("Failed to load Public key %w", err)
	}
	pubKey, err := age.ParseX25519Recipient(string(keyByte))
	if err != nil {
		return nil, fmt.Errorf("Failed to load Public key %w", err)
	}

	return pubKey, nil
}

func EncryptData(input io.ReadCloser, recipent age.Recipient) ([]byte, error) {
	encryptedData := &bytes.Buffer{}
	dec, err := age.Encrypt(encryptedData, recipent)
	if err != nil {
		return nil, fmt.Errorf("Failed to encrypt data, can't create encrypter, %v", err)
	}
	defer dec.Close()
	var buffer []byte = make([]byte, 4096)
	for {
		var n int
		n, err = io.ReadFull(input, buffer)
		if n == 0 {
			break
		}
		dec.Write(buffer[:n])
	}
	if !errors.Is(err, io.EOF) && !errors.Is(err, io.ErrUnexpectedEOF) {
		return nil, fmt.Errorf("Failed to encrypt data, can't read input, %v", err)
	}
	if dec.Close() != nil {
		return nil, fmt.Errorf("Failed to encrypt data, can't close encrypter, %v", err)
	}
	return encryptedData.Bytes(), nil
}

func DencryptData(input io.ReadCloser, identity age.Identity) ([]byte, error) {
	out, err := age.Decrypt(input, identity)
	if err != nil {
		return nil, fmt.Errorf("Failed to decrypt data %w", err)
	}
	text, err := io.ReadAll(out)
	if err != nil {
		return nil, fmt.Errorf("Failed to read decrypt data %w", err)
	}
	return text, nil
}
