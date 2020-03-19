// https://myaccount.google.com/lesssecureapps

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

	flag.StringVar(&from, "f", "", "sender of the mail")
	flag.StringVar(&message, "m", "no message", "content of the mail")
	flag.StringVar(&pass, "p", "", "password to connect to your email account")
	flag.StringVar(&subject, "s", "no subject", "subject of the mail")
	flag.StringVar(&recipient, "t", "", "recipient email address")

	flag.Parse()

	fmt.Println("flags parsed")
	if recipient == "" || from == "" || pass == "" {
		fmt.Println("You need to provide correct recipient, sender email addresses and a correct password")
		os.Exit(0)
	}

	var (
		//recipients = []string{recipient}
		msg = []byte("From: " + from + "\n" +
			"To: " + recipient + "\n" +
			"Subject: " + subject + "\n\n" +
			message)
	)

	auth := smtp.PlainAuth("", from, pass, "smtp.gmail.com")
	fmt.Printf("plain auth : success\n")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{recipient}, msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("send mail : success\n")
}
