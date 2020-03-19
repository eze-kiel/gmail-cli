package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/namsral/flag"
)

func main() {
	var recipient, subject, message, from, pass string

	// Declaration of flags
	flag.StringVar(&from, "f", "", "sender of the mail")
	flag.StringVar(&message, "m", "no message", "content of the mail")
	flag.StringVar(&pass, "p", "", "password to connect to your email account")
	flag.StringVar(&subject, "s", "no subject", "subject of the mail")
	flag.StringVar(&recipient, "t", "", "recipient email address")

	flag.Parse()

	fmt.Println("flags parsed")

	// This statement check if the principal flags are used
	// If not, the program exit
	if recipient == "" || from == "" {
		fmt.Println("You need to provide correct recipient, sender email addresses and a correct password")
		os.Exit(0)
	}

	msg := []byte("From: " + from + "\n" +
		"To: " + recipient + "\n" +
		"Subject: " + subject + "\n\n" +
		message)

	// SMTP authentification to the GMail server
	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")

	// Port 587 is used because it uses TLS
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{recipient}, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("send mail : success\n")
}
