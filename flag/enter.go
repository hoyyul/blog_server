package flag

import sys_flag "flag"

type Option struct {
	Version bool
	DB      bool
}

func Parse() Option {
	version := sys_flag.Bool("v", false, "version")
	db := sys_flag.Bool("db", false, "Initialize database")

	sys_flag.Parse()
	return Option{
		*version,
		*db,
	}
}

func IsWebStop(op Option) bool {
	return op.DB
}

func RunOption(op Option) {
	if op.DB {
		MigrateTables()
	}
}
