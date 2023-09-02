package pwd

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPwd same pwd will not share same hash value
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
