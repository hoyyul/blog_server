package ctype

import "encoding/json"

type Role int

const (
	PermissionAdmin       Role = 1
	PermissionUser        Role = 2
	PermissionVistor      Role = 3
	PermissionDisableUser Role = 4
)

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r Role) String() string {
	var str string
	switch r {
	case PermissionAdmin:
		str = "Adminstrator"
	case PermissionUser:
		str = "User"
	case PermissionVistor:
		str = "Vistor"
	case PermissionDisableUser:
		str = "DisableUser"
	default:
		str = "Other"
	}
	return str
}
