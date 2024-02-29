package hasher

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Hash(pwd string) string {
	bpwd := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(bpwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
