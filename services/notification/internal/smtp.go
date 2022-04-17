package internal

import (
	"log"
	"net"
	"net/smtp"

	"github.com/jonashiltl/sessions-backend/services/notification/internal/config"
)

func Connect(c config.Config) (*smtp.Client, error) {
	// auth := smtp.PlainAuth("", c.SMTP_USERNAME, c.SMTP_PORT, c.SMTP_HOST)

	// Connect to the server using its address
	conn, err := net.Dial("tcp", c.SMTP_HOST+":"+c.SMTP_PORT)
	if err != nil {
		log.Print("Conn Error:")
		return nil, err
	}

	client, err := smtp.NewClient(conn, c.SMTP_HOST)
	if err != nil {
		log.Print("Client Error:")
		return nil, err
	}

	/* 	if err = client.StartTLS(&tls.Config{}); err != nil {
	   		log.Print("TLS Error:")
	   		return nil, err
	   	}
	*/

	/* if err = client.Auth(auth); err != nil {
		log.Print("Auth Error:")
		return nil, err
	} */

	return client, nil
}
