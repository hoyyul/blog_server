package flag

import (
	sys_flag "flag"

	"github.com/fatih/structs"
)

type Option struct {
	DB   bool
	User string
	ES   bool
	Dump string
	Load string
}

func Parse() Option {
	db := sys_flag.Bool("db", false, "Initialize database")
	user := sys_flag.String("u", "", "Create user")
	es := sys_flag.Bool("es", false, "es operation")
	dump := sys_flag.String("dump", "", "dump index to json")
	load := sys_flag.String("load", "", "load json to index")

	sys_flag.Parse()
	return Option{
		*db,
		*user,
		*es,
		*dump,
		*load,
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
	if op.ES {
		//global.ESClient = initialization.EsConnect()
		//EsCreateIndex()
		if op.Dump != "" {
			DumpIndex(op.Dump)
		}
		if op.Load != "" {
			LoadIndex(op.Load)
		}
		return
	}
	sys_flag.Usage()
}
