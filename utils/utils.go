package utils

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"

	"github.com/go-playground/validator"
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

func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, err := range errs {
			if f, ok := getObj.Elem().FieldByName(err.Field()); ok { //find the first error appeared in fieldname
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}
