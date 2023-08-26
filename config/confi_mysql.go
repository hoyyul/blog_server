package config

import "strconv"

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Config   string `yaml:"config"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

func (m Mysql) DSN() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DB + "?" + m.Config

}
