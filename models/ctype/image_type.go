package ctype

import "encoding/json"

type StorageLocation int

const (
	Local StorageLocation = 1
	AWS   StorageLocation = 2
)

func (s StorageLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s StorageLocation) String() string {
	var str string
	switch s {
	case Local:
		str = "QQ"
	default:
		str = "Other"
	}
	return str
}
