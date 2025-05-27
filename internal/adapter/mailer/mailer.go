package mailer

import (
	"context"
	"fmt"
	"github.com/mailersend/mailersend-go"
	"github.com/sorawaslocked/ap2final_email_service/internal/model"
	"log"
	"time"
)

type Mailer struct {
	client *mailersend.Mailersend
}

func NewMailer(client *mailersend.Mailersend) *Mailer {
	return &Mailer{client: client}
}

func (m *Mailer) Send(ctx context.Context, user model.User) error {
	ctxC, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	subject := fmt.Sprintf("Welcome, %s", user.Email)
	text := "Greetings! You got this message from MailerSend."
	html := "Greetings! You got this message from MailerSend."

	from := mailersend.From{
		Name:  "Movie Service",
		Email: "MS_wvWhOk@test-z0vklo6nqwxl7qrx.mlsender.net",
	}

	recipients := []mailersend.Recipient{
		{
			Email: user.Email,
		},
	}

	message := m.client.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetText(text)

	res, err := m.client.Email.Send(ctxC, message)
	if err != nil {
		log.Println("m.client.Email.Send:", err, "Status:", res.Status, "Customer:", user.Email)
		log.Println("Try to send email again to: ", user.Email)
		_, err = m.client.Email.Send(ctxC, message)
	}

	return err
}
