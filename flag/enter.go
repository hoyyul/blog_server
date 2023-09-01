package flag

import sys_flag "flag"

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

func IsWebStop(op Option) bool {
	return op.DB || op.User == "user" || op.User == "admin"
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
}
