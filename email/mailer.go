// Package email includes all functions required for sending emails using golang
package email

import (
	"github.com/programmer-richa/utility/constants"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/smtp"
)

// Mime type supported by email
const (
	// For sending HTML content (Passed as an argument to send function)
	MIME_HTML = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

// MailRequest represents struct to map email receiver(s), subject and body
type MailRequest struct {
	from    string
	to      []string
	subject string
	body    string
	// The following parameters are for mail server authentication
	server   string
	port     int
	username string
	password string
	// Template pointer to parse template files
	tpl *template.Template
}

// NewMailRequest returns pointer to MailRequest
func NewMailRequest(to []string, subject string, server string, port int, username string, password string, tpl *template.Template) *MailRequest {
	return &MailRequest{
		to:       to,
		subject:  subject,
		server:   server,
		port:     port,
		username: username,
		password: password,
		tpl:      tpl,
	}
}

// ParseTemplate binds the data passed with the email template file.
func (r *MailRequest) ParseTemplate(fileName string, data interface{}) error {
	if r.tpl == nil {
		// Return custom error
		return errors.New(constants.NilTpl)
	}
	buffer := new(bytes.Buffer)
	if err := r.tpl.ExecuteTemplate(buffer, fileName, data); err != nil {
		return err
	}
	r.body = buffer.String()
	return nil
}

// SendMail sends the email and returns true if email is sent successfully
func (r *MailRequest) SendMail(mime string) error {
	body := "To: " + r.to[0] + "\r\nSubject: " + r.subject + "\r\n" + mime + "\r\n" + r.body
	SMTP := fmt.Sprintf("%s:%d", r.server, r.port)
	err := smtp.SendMail(SMTP, smtp.PlainAuth("", r.username, r.password, r.server), r.username, r.to, []byte(body))
	return err
}

// Send sends the email and saves the result in the log file
func (r *MailRequest) Send(templateName string, mime string, items interface{}) (error, bool) {
	err := r.ParseTemplate(templateName, items)
	if err != nil {
		return err, false
	}
	err = r.SendMail(mime)
	return err, err == nil

}
