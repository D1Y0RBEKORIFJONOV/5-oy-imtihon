package email

import (
	"ekzamen_5/notification-service/internal/entity"
	"encoding/json"
	"fmt"
	"net/smtp"
)

func SenSecretCode(email entity.EmailNotificationReq) error {
	from := "yusupovabduazim0811@gmail.com"
	password := "slws nzfk namk eali"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: Your Secret Code\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	reqByte, err := json.Marshal(&email)
	if err != nil {
		return err
	}
	body := fmt.Sprintf("%s", reqByte)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, email.Recipient, message)
	if err != nil {
		return err
	}
	return nil
}
