package helpers

import (
	"email-service/model"
	"fmt"
	"net/smtp"
	"os"
)

// Send a verification email
func SendVerificationEmail(data model.UserCredential) error {
	authEmail := os.Getenv("AUTH_EMAIL")
	authPass := os.Getenv("AUTH_PASS")
	smptHost := os.Getenv("SMPT_HOST")
	smptPort := os.Getenv("SMPT_PORT")
	
	url := fmt.Sprintf("https://carstruck-4d6b89ee5e4e.herokuapp.com/users/verify/%v/%v", data.Id, data.Token)
	subject := "Subject: Jeki Account Verification\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body, err := VerificationEmailBody(data.FullName, url)
	if err != nil {
		return err
	}

	smptAddr := fmt.Sprintf("%s:%s", smptHost, smptPort)
	smptAuth := smtp.PlainAuth("", authEmail, authPass, smptHost)
	msg := []byte(subject + mime + body)

	err = smtp.SendMail(smptAddr, smptAuth, authEmail, []string{data.Email}, msg)
	if err != nil {
		return err
	}
	
	fmt.Printf("Verification email sent to %v\n", data.Email)
	return nil
}