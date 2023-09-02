package flag

import (
	sys_flag "flag"

	"github.com/fatih/structs"
)

type Option struct {
	Version bool
	DB      bool
	User    string
}

func Parse() Option {
	version := sys_flag.Bool("v", false, "version")
	db := sys_flag.Bool("db", false, "Initialize database")
	user := sys_flag.String("u", "", "Create user")

	sys_flag.Parse()
	return Option{
		*version,
		*db,
		*user,
	}
}

func IsWebStop(option Option) (b bool) {
	maps := structs.Map(&option)
	for _, v := range maps {
		switch val := v.(type) {
		case string:
			if val != "" {
				b = true
			}
		case bool:
			if val {
				b = true
			}
		}
	}
	return b
}

func RunOption(op Option) {
	if op.DB {
		MigrateTables()
		return
	}
	if op.User == "user" || op.User == "admin" {
		CreateUser(op.User)
		return
	}
	sys_flag.Usage()
}
