package secrets

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"github.com/pkg/errors"
	"os"
	"planetbi/db"
	"time"
)

var secretKey = ""

func Init() {
	// get secrets from environment variables
	secretKey = os.Getenv("SECRET_KEY")
	if secretKey == "" {
		panic("SECRET_KEY not set")
	}
}

// https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d#code
func encrypt(msg string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	// We need a 12-byte nonce for GCM (modifiable if you use cipher.NewGCMWithNonceSize())
	// A nonce should always be randomly generated for every encryption.
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	// ciphertext here is actually nonce+ciphertext
	// So that when we decrypt, just knowing the nonce size
	// is enough to separate it from the ciphertext.
	encryptedText := gcm.Seal(nonce, nonce, []byte(msg), nil)

	return base64.StdEncoding.EncodeToString(encryptedText), nil
}

// https://dev.to/breda/secret-key-encryption-with-go-using-aes-316d#code
func decrypt(encryptedText string) (string, error) {
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	encryptedTextBytes, err := base64.StdEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	encryptedText = string(encryptedTextBytes)

	// Since we know the encryptedText is actually nonce+encryptedText
	// And len(nonce) == NonceSize(). We can separate the two.
	nonceSize := gcm.NonceSize()

	nonce, encryptedText := encryptedText[:nonceSize], encryptedText[nonceSize:]

	msg, err := gcm.Open(nil, []byte(nonce), []byte(encryptedText), nil)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	return string(msg), nil
}

func Set(ctx context.Context, id, secret string) error {
	value, err := encrypt(secret)
	if err != nil {
		return errors.Wrap(err, "")
	}
	_, err = db.Pool.Exec(ctx, "INSERT INTO secrets(id, value, updated_at) VALUES($1, $2, $3) ON CONFLICT (id) DO UPDATE SET value = $2, updated_at = $3", id, value, time.Now())
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func Get(ctx context.Context, id string) (string, error) {
	var value string
	err := db.Pool.QueryRow(ctx, "SELECT value FROM secrets WHERE id=$1", id).Scan(&value)
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	secret, err := decrypt(value)
	if err != nil {
		return "", errors.Wrap(err, "")
	}
	return secret, nil
}

func Delete(ctx context.Context, id string) error {
	_, err := db.Pool.Exec(ctx, "DELETE FROM secrets WHERE id=$1", id)
	if err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}

func IsNotFoundErr(err error) bool {
	if errors.Is(err, db.ErrNoRows) {
		return true
	}

	return false
}
