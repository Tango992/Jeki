package helper

import (
	"email-service/model"
	"fmt"
	"log"
	"net/smtp"
	"os"
)

// Send a verification email
func SendVerificationEmail(data model.UserCredential) error {
	verificationUrl := os.Getenv("VERIFICATION_URL")
	authEmail := os.Getenv("AUTH_EMAIL")
	authPass := os.Getenv("AUTH_PASS")
	smptHost := os.Getenv("SMPT_HOST")
	smptPort := os.Getenv("SMPT_PORT")
	
	url := fmt.Sprintf("%v/%v/%v", verificationUrl, data.Id, data.Token)
	fmt.Println(url)
	subject := "Subject: Jeki Account Verification\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	body, err := VerificationEmailBody(data.Name, url)
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
	
	log.Printf("Verification email sent to %v\n", data.Email)
	return nil
}