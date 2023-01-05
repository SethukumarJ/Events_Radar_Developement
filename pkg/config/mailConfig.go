package config

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
)

type MailConfig interface {
	SendMail(cfg Config, to string, message string) error
}

type mailConfig struct{}

func NewMailConfig() MailConfig {
	return &mailConfig{}
}

func (c *mailConfig) SendMail(cfg Config, to string, message string) error {

	fmt.Printf("\n\nemail :  %v\n\n", to)
	log.Println("Email Id to send message : ", to)
	userName := cfg.SMTPUSERNAME
	password := cfg.SMTPPASSWORD
	smtpHost := cfg.SMTPHOST
	smtpPort := cfg.SMTPPORT

	auth := smtp.PlainAuth("", userName, password, smtpHost)
	fmt.Printf("\n\nauth : %v\n\n", auth)

	headers := make(map[string]string)

	headers["Subject"] = "Events Radar"
	headers["From"] = userName

	var msg bytes.Buffer
	for k, v := range headers {
		msg.WriteString(k + ": " + v + "\r\n")
	}

	msg.WriteString("\r\n")
	msg.WriteString(message)

	// sending email
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, userName, []string{to}, msg.Bytes())

}
