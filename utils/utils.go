package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func InList(key string, list []string) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}
