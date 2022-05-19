// Sending Email Using Smtp in Golang
package main

import (
	"fmt"
	"log"

	mail "github.com/xhit/go-simple-mail"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

// Main function
func main() {

	server := mail.NewSMTPClient()
	server.Host = "smtp.mailtrap.io"
	server.Port = 2525
	server.Username = "275a54a5db8a21"
	server.Password = "f49cb18f409064"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("From Me <me@host.com>")
	email.AddTo("you@example.com")
	email.AddCc("another_you@example.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully sent mail to all user in toList")
}
