package utils

import (
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendMail(from, to, subject, body string) (ok bool, err error) {
	mail := gomail.NewMessage()
	mail.SetHeader("From", from)
	mail.SetHeader("To", to)
	//mail.SetAddressHeader("Cc", "dan@example.com", "Dan")
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", body)

	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(mail); err != nil {
		return false, err
	}
	return true, nil
}
