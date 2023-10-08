package utils

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	if errs, ok := err.(validator.ValidationErrors); ok {

		if getObj.Kind() != reflect.Ptr {
			return "Invalid object type. Expected a pointer."
		}

		for _, err := range errs {
			if f, ok := getObj.Elem().FieldByName(err.Field()); ok {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return err.Error()
}
