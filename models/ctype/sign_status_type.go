package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ     SignStatus = 1
	SignGithub SignStatus = 2
	SignEmail  SignStatus = 3
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGithub:
		str = "Github"
	case SignEmail:
		str = "Email"
	default:
		str = "Other"
	}
	return str
}
