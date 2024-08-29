package email

import (
	"ekzamen_5/notification-service/internal/entity"
	"fmt"
	"net/smtp"
)

func SendSecretCode(email *entity.EmailNotificationReq) error {
	from := "diyordev3@gmail.com"
	password := "uvcy ksvb yxyy gsmi"

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Subject: " + email.Tittle + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Email Notification</title>
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
			<style>
				body {
					font-family: Arial, sans-serif;
					color: #ffffff;
					background-color: #0a0a0a;
					margin: 0;
					padding: 0;
				}
				.container {
					max-width: 600px;
					margin: 20px auto;
					padding: 40px;
					background-color: #1e1e1e;
					border-radius: 12px;
					box-shadow: 0 5px 20px rgba(0, 0, 0, 0.3);
					border: 1px solid #333;
				}
				.title {
					color: #4caf50;
					font-size: 32px;
					font-weight: bold;
					text-align: center;
					margin-bottom: 30px;
				}
				.title i {
					margin-right: 10px;
				}
				.content {
					font-size: 20px;
					line-height: 1.8;
					color: #e0e0e0;
					margin-bottom: 40px;
					text-align: center;
				}
				.footer {
					margin-top: 40px;
					font-size: 16px;
					color: #bdbdbd;
					text-align: center;
				}
				hr {
					border: none;
					border-top: 1px solid #444;
					margin: 40px 0;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<h1 class="title">
					<i class="fas fa-bell"></i> %s
				</h1>
				<p class="content">%s</p>
				<hr />
				<div class="footer">
					<p>Отправлено: %s</p>
					<p>С уважением, <span style="color: #4caf50;">%s</span></p>
				</div>
			</div>
		</body>
		</html>`,
		email.Tittle,
		email.Content,
		email.SenderAt.Format("02.01.2006 15:04"),
		email.SenderName,
	)

	message := []byte(subject + mime + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, email.Recipient, message)
	if err != nil {
		return err
	}

	return nil
}
