package helpers

import (
	"bytes"
	"crypto/rand"
	"io"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/scrypt"
)

// Changed the c.MustBindWith() ->  c.ShouldBindWith().
// I don't want to auto return 400 when error happened.
// origin function is here: https://github.com/gin-gonic/gin/blob/master/context.go
func Bind(c *gin.Context, obj interface{}) error {
	// Request method is: post, get...
	// content type can be json, yaml,...
	b := binding.Default(c.Request.Method, c.ContentType())

	// shouldBingWith can notify the developer for approprisate modification
	return c.ShouldBindWith(obj, b)
}

func GenToken(id uint) string {
	jwt_token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwt_token.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwt_token.SignedString([]byte(NBSecretPassword))
	return token
}

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
	// log.Printf("New Satl is created for %v is %x \n", username, salt)
	if err != nil {
		log.Println(err)
		return []byte{}, []byte{}, err
	}
	// log.Printf("Pwd to cal Hash %v ", password)
	hash, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
		return []byte{}, []byte{}, err
	}
	// log.Printf("New Hash is created for %v is %x \n", username, hash)
	return salt, hash, nil
}

func VerifyHashWithSalt(password string, salt []byte, hash []byte) bool {
	// log.Printf("Salt for verify %x \n ", salt)
	// log.Printf("Pwd to verify Hash %v ", password)
	res, err := scrypt.Key([]byte(password), salt, 1<<14, 8, 1, PW_HASH_BYTES)
	if err != nil {
		log.Fatal(err)
		return false
	}
	// log.Printf("Calculated Hash %x \n Stored Hash %x", res, hash)
	if bytes.Compare(hash, res) == 0 {
		return true
	}
	return false
}
