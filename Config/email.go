package config

import "os"

type SMTPConfig struct {
	Host string
	Port int
	User string
	Pass string
	From string
}

//MUST start with CAPITAL L
func LoadSMTPConfig() SMTPConfig {
	return SMTPConfig{
		Host: os.Getenv("SMTP_HOST"),
		Port: 587,
		User: os.Getenv("SMTP_USER"),
		Pass: os.Getenv("SMTP_PASS"),
		From: os.Getenv("FROM_EMAIL"),
	}
}
