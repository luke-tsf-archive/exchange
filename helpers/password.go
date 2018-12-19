package helpers

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

const (
	PW_SALT_BYTES = 32
	PW_HASH_BYTES = 64
)

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

func GenerateSaltAndHash(username string, password string) (string, string, error) {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	log.Printf("New Satl is created for %v is %x", username, salt)
	if err != nil {
		log.Println(err)
		return "", "", err
	}
	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}
	log.Printf("New Hash is created for %v is %x", username, hash)
	fmt.Printf("%x\n", hash)
	return string(salt), string(hash), nil
}

func VerifyHashWithSalt(password string, salt string, hash string) bool {
	res, err := scrypt.Key([]byte(password), []byte(salt), 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if bytes.Compare([]byte(hash), res) == 0 {
		return true
	}
	return false
}
