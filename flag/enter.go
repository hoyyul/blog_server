package flag

import (
	sys_flag "flag"

	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string
	ES   string
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "Initialize database")
	user := sys_flag.String("u", "", "Create user")
	es := sys_flag.String("es", "", "es operation")

	sys_flag.Parse()
	return Option{
		*db,
		*user,
		*es,
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
	if op.ES == "create" {
		EsCreateIndex()
		return
	}
	sys_flag.Usage()
}
