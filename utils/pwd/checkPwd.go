package pwd

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CheckPwd(hashPwd string, pwd string) bool {
	byteHash := []byte(hashPwd)

	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
