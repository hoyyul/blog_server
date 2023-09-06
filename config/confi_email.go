package config

type Email struct {
	Host        string `json:"host" yaml:"host"`
	Port        int    `json:"port" yaml:"port"`
	SenderEmail string `json:"sender_email" yaml:"sender_email"`
	Password    string `json:"password" yaml:"password"`
	SenderName  string `json:"sender_name" yaml:"sender_name"`
	UseSSL      bool   `json:"use_ssl" yaml:"use_ssl"`
	UserTls     bool   `json:"user_tls" yaml:"user_tls"`
}
