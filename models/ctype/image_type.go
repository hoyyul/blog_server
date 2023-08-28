package ctype

import "encoding/json"

type StorageLocatioin int

const (
	Local StorageLocatioin = 1
)

func (s StorageLocatioin) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s StorageLocatioin) String() string {
	var str string
	switch s {
	case Local:
		str = "QQ"
	default:
		str = "Other"
	}
	return str
}
