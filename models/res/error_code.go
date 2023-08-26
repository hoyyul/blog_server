package res

type ErrorCode int

const (
	SettingsError  = 1001
	ParameterError = 1002
)

var (
	ErrorMap = map[ErrorCode]string{
		SettingsError:  "Settings Error",
		ParameterError: "Parameter Error",
	}
)
