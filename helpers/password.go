package helpers

import (
	"bytes"
	"crypto/rand"
	"io"
	"log"

	"golang.org/x/crypto/scrypt"
)

const (
	PW_SALT_BYTES    = 32
	PW_HASH_BYTES    = 64
	NBRandomPassword = "A String Very Very Very Niubilty!!@##$!@#4"
)

func IsEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}

func GenerateSaltAndHash(username string, password string) ([]byte, []byte, error) {
	salt := make([]byte, PW_SALT_BYTES)
	_, err := io.ReadFull(rand.Reader, salt)
	log.Printf("New Satl is created for %v is %x", username, salt)
	if err != nil {
		log.Println(err)
		return []byte{}, []byte{}, err
	}
	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
		return []byte{}, []byte{}, err
	}
	log.Printf("New Hash is created for %v is %x", username, hash)
	return salt, hash, nil
}

func VerifyHashWithSalt(password string, salt []byte, hash []byte) bool {
	res, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if bytes.Compare(hash, res) == 0 {
		return true
	}
	return false
}
