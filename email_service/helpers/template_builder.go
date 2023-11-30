package helpers

import (
	"bytes"
	"html/template"
)

// Inserts user data into the HTML template
func VerificationEmailBody(name, url string) (string, error) {
	tmpl, err := template.ParseFiles("./template/verification_email.html")
	if err != nil {
		return "", err
	}
	
	templateData := map[string]string{
		"Name": name,
		"URL": url,
	}
	
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, templateData); err != nil {
		return "", err
	}
	return buf.String(), nil
}