package main

import (
	"fmt"
	"github.com/anonymfrominternet/Hotel/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
	"time"
)

func listenForMail() {
	go func() {
		for {
			message := <-appConfig.MailChan
			sendMessage(message)
		}
	}()
}

func sendMessage(mailData models.MailData) {
	// Creating and configuring the server
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 15 * time.Second
	server.SendTimeout = 15 * time.Second

	client, err := server.Connect()
	if err != nil {
		fmt.Println("cannot get SMTPClient from server.Connect()")
	}

	email := mail.NewMSG()
	email.SetFrom(mailData.From).AddTo(mailData.To).SetSubject(mailData.Subject)
	email.SetBody(mail.TextPlain, mailData.Content)

	err = email.Send(client)
	if err != nil {
		fmt.Println("cannot send an email")
	}
}
