package utils

import (
	"reflect"

	"github.com/go-playground/validator"
)

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
