package notification

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"

	"github.com/ordomigato/parking-app/initializers"
)

type EmailMessage struct {
	From          string
	To            []string
	Subject       string
	TemplatePath  string
	PlainTextPath string
	Data          map[string]interface{}
}

func SendEMail(config *initializers.Config, msg EmailMessage) {
	fmt.Println("SENDING EMAIL")
	auth := smtp.PlainAuth(
		"",
		config.MailUsername,
		config.MailPassword,
		config.MailHost,
	)

	t, err := template.ParseFiles(msg.TemplatePath)
	if err != nil {
		fmt.Println(err.Error())
	}

	var body bytes.Buffer

	t.Execute(&body, msg.Data)

	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"

	m := "Subject: " + msg.Subject + "\n" + headers + body.String()

	err = smtp.SendMail(
		config.MailHost+":"+config.MailPort,
		auth,
		msg.From,
		msg.To,
		[]byte(m),
	)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("EMAIL SENT")

}
