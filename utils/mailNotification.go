package utils

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"strconv"
	"os"
	"path/filepath"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL string
	FirstName string
	Subject string
	MailTo string
}


// 👇 Email template parser
func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return template.ParseFiles(paths...)
}

// ? Email template parser
func SendEmail(data *EmailData) error {
	// Sender data.
	from := GetEnvVar("EMAIL_FROM")
	smtpPass := GetEnvVar("SMTP_PASS")
	smtpUser := GetEnvVar("SMTP_USER")
	smtpHost := GetEnvVar("SMTP_HOST")
	portString := GetEnvVar("SMTP_PORT")

	smtpPort, err := strconv.Atoi(portString)
	if err != nil {
        // ... handle error
        log.Fatal("Could not covert port string", err)
    }

	var body bytes.Buffer
	template, err := ParseTemplateDir("views/templates")
	if err != nil {
		log.Fatal("Could not parse template", err)
	}

	template.ExecuteTemplate(&body, "verificationCode.html", &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", data.MailTo)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}