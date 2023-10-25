package helpers

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(code int, mail string) bool {
	d := gomail.NewDialer(os.Getenv("MAIL_PROTOCOL"), 587, os.Getenv("MAIL_USER"), os.Getenv("MAIL_PASSWORD"))

	m := gomail.NewMessage()
	m.SetHeader("From", "leomag.sil@gmail.com")
	m.SetHeader("To", mail)
	m.SetHeader("Subject", "Código de verificação")
	m.SetBody("text/plain", "Olá, aqui está seu código de verificação: \n\n"+strconv.Itoa(code))

	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err.Error())
		return false
	}

	return true
}
