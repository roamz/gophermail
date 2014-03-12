package gophermail

import (
	"net/smtp"
)

// SendMail connects to the server at addr, switches to TLS if possible,
// authenticates with mechanism a if possible, and then sends the given Message.
//
// Based heavily on smtp.SendMail().
func SendMail(addr string, a smtp.Auth, msg *Message) error {
	fn, err := msg.SendMailArgs(addr, a)
	if err != nil {
		return err
	}

	return smtp.SendMail(fn())
}
