package email

import (
	"github.com/programmer-richa/utility/constants"
	"fmt"
	"html/template"
	"testing"
)

// TestMailer runs several test cases to check the correctness of
//the Mail functionality defined in functions package.
func TestMailer(t *testing.T) {
	username := constants.MailUsername
	password := constants.MailPassword
	server := constants.MailServer
	port := constants.MailPort
	data := map[string]string{"URL": "http://abc.com",
		"Name":         "Richa",
		"Designer":     "Programmer Richa",
		"DesignerSite": "http://abcd.com",
	}

	tpl := template.Must(template.New("").ParseGlob("../templates/emails/*.gohtml"))
	request1 := NewMailRequest([]string{"programmer.richa@gmail.com"},
		"Test Mail",
		server, port, username, password, nil)

	request2 := NewMailRequest([]string{"programmer.richa@gmail.com"},
		"Test Mail",
		server, port, username, password, tpl)

	request3 := NewMailRequest([]string{"programmer.richa@gmail.com"},
		"Test Mail",
		server, port, username+"invalid", password, tpl)

	tests := []struct {
		name     string
		data     map[string]string
		mail     *MailRequest
		filename string
		valid    bool
	}{
		{
			"nil template pointer for email parsing",
			data,
			request1,
			"account_verification.gohtml",
			false,
		},
		{
			"Valid template pointer for email parsing",
			data,
			request2,
			"account_verification.gohtml",
			true,
		},
		{
			"Invalid mail authentication credentials",
			data,
			request3,
			"account_verification.gohtml",
			false,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err, result := c.mail.Send(c.filename, MIME_HTML, data)
			if result != c.valid {
				t.Fatal("Mailer Function ", c.name, result, err)
			} else {
				fmt.Println("Mailer Function-", c.name, "Pass")
			}
		})
	}
}
